package memeCoinCtrl

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"meme_coin_api/service/internal/model"
	boMemeCoin "meme_coin_api/service/internal/model/bo/memeCoin"
	"meme_coin_api/service/internal/utils"
	redisDao "meme_coin_api/service/repository/caches/redis"
	"meme_coin_api/service/repository/orm/gormDao"
)

type MemeCoinCtrl interface {
	Create(ctx context.Context, args boMemeCoin.CreateArgs) (boMemeCoin.CreateReply, error)
	Get(ctx context.Context, args boMemeCoin.GetArgs) (boMemeCoin.GetReply, error)
	Update(ctx context.Context, args boMemeCoin.UpdateArgs) error
	IncreasePopularityScore(ctx context.Context, args boMemeCoin.IncreasePopularityScoreArgs) error
	Delete(ctx context.Context, args boMemeCoin.DeleteArgs) error
}

func New(pack memeCoinCtrlPack) MemeCoinCtrl {
	return &memeCoinCtrl{
		pack: pack,
	}
}

type memeCoinCtrlPack struct {
	dig.In

	RedisMemeCoin *redis.Client `name:"meme_coin"`
	MySQLMemeCoin *gorm.DB      `name:"meme_coin"`
}

type memeCoinCtrl struct {
	pack memeCoinCtrlPack
}

func (ctrl *memeCoinCtrl) Create(ctx context.Context, args boMemeCoin.CreateArgs) (boMemeCoin.CreateReply, error) {
	var reply boMemeCoin.CreateReply
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	id := utils.GetSnowflakeIDInt64()
	coinInfo := &model.CoinInfo{
		ID:          id,
		Name:        args.Name,
		Description: args.Description,
	}
	if err := dao.Create(coinInfo); err != nil {
		return reply, err
	}
	reply.ID = id
	return reply, nil
}

func (ctrl *memeCoinCtrl) Get(ctx context.Context, args boMemeCoin.GetArgs) (boMemeCoin.GetReply, error) {
	var reply boMemeCoin.GetReply
	caches := redisDao.NewMemeCoinCache(ctrl.pack.RedisMemeCoin)
	coin, err := caches.Get(ctx, args.ID)
	if err != nil {
		return reply, err
	}
	if coin != nil {
		reply.MemeCoin = coin
		return reply, nil
	}

	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	coin, err = dao.Get(args.ID)
	if err != nil {
		return reply, err
	}
	caches.Set(ctx, coin)
	reply.MemeCoin = coin
	return reply, nil
}

func (ctrl *memeCoinCtrl) Update(ctx context.Context, args boMemeCoin.UpdateArgs) error {
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	if err := dao.UpdateDescription(args.ID, args.Description); err != nil {
		return err
	}
	caches := redisDao.NewMemeCoinCache(ctrl.pack.RedisMemeCoin)
	caches.Delete(ctx, args.ID)
	return nil
}

func (ctrl *memeCoinCtrl) IncreasePopularityScore(ctx context.Context, args boMemeCoin.IncreasePopularityScoreArgs) error {
	// TODO: Utilize a mq and cronjob to enhance availability and scalability.
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	if err := dao.IncreasePopularityScore(args.ID); err != nil {
		return err
	}
	caches := redisDao.NewMemeCoinCache(ctrl.pack.RedisMemeCoin)
	caches.Delete(ctx, args.ID)
	return nil
}

func (ctrl *memeCoinCtrl) Delete(ctx context.Context, args boMemeCoin.DeleteArgs) error {
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	if err := dao.DeleteMemeCoin(args.ID); err != nil {
		return err
	}
	caches := redisDao.NewMemeCoinCache(ctrl.pack.RedisMemeCoin)
	return caches.Delete(ctx, args.ID)
}
