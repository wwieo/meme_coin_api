package memeCoinCtrl

import (
	"context"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"meme_coin_api/service/dao/orm/gormDao"
	"meme_coin_api/service/internal/errorx"
	"meme_coin_api/service/internal/model"
	boMemeCoin "meme_coin_api/service/internal/model/bo/memeCoin"
	"meme_coin_api/service/internal/utils"
)

type MemeCoinCtrl interface {
	Create(ctx context.Context, args boMemeCoin.CreateArgs) error
	Get(ctx context.Context, args boMemeCoin.GetArgs) (boMemeCoin.GetReply, error)
	Update(ctx context.Context, args boMemeCoin.UpdateArgs) error
	IncreasePopularityScore(ctx context.Context, args boMemeCoin.IncreasePopularityScoreArgs) error
	DeleteMemeCoin(ctx context.Context, args boMemeCoin.DeleteArgs) error
}

func New(pack memeCoinCtrlPack) MemeCoinCtrl {
	return &memeCoinCtrl{
		pack: pack,
	}
}

type memeCoinCtrlPack struct {
	dig.In

	MySQLMemeCoin *gorm.DB `name:"meme_coin"`
}

type memeCoinCtrl struct {
	pack memeCoinCtrlPack
}

func (ctrl *memeCoinCtrl) Create(ctx context.Context, args boMemeCoin.CreateArgs) error {
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	if exist, err := dao.Exist(args.Name); err != nil || exist {
		return errorx.RecordExisted
	}

	id := utils.GetSnowflakeIDInt64()
	coinInfo := &model.CoinInfo{
		ID:          id,
		Name:        args.Name,
		Description: args.Description,
	}
	return dao.Create(coinInfo)
}

func (ctrl *memeCoinCtrl) Get(ctx context.Context, args boMemeCoin.GetArgs) (boMemeCoin.GetReply, error) {
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	memeCoin, err := dao.Get(args.ID)
	if err != nil {
		return boMemeCoin.GetReply{}, err
	}
	return boMemeCoin.GetReply{MemeCoin: memeCoin}, nil
}

func (ctrl *memeCoinCtrl) Update(ctx context.Context, args boMemeCoin.UpdateArgs) error {
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	return dao.UpdateDescription(args.ID, args.Description)
}

func (ctrl *memeCoinCtrl) IncreasePopularityScore(ctx context.Context, args boMemeCoin.IncreasePopularityScoreArgs) error {
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	return dao.IncreasePopularityScore(args.ID)
}

func (ctrl *memeCoinCtrl) DeleteMemeCoin(ctx context.Context, args boMemeCoin.DeleteArgs) error {
	dao := gormDao.NewMemeCoinDao(ctrl.pack.MySQLMemeCoin)
	return dao.DeleteMemeCoin(args.ID)
}
