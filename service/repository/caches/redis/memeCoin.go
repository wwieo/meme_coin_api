package redisDao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"meme_coin_api/service/internal/model"
	"meme_coin_api/service/repository/caches"
	"strconv"
	"time"
)

func NewMemeCoinCache(r *redis.Client) caches.MemeCoinCache {
	return &memeCoinCache{
		prefixInfoKey: "meme_coin:info",
		client:        r,
	}
}

type memeCoinCache struct {
	prefixInfoKey string
	client        *redis.Client
}

func (dao *memeCoinCache) buildInfoKey(id int64) string {
	return fmt.Sprintf("%s:%d", dao.prefixInfoKey, id)
}

func (dao *memeCoinCache) Exist(ctx context.Context, id int64) (bool, error) {
	result, err := dao.client.Exists(ctx, dao.buildInfoKey(id)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (dao *memeCoinCache) Get(ctx context.Context, id int64) (*model.MemeCoin, error) {
	key := dao.buildInfoKey(id)
	values, err := dao.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	// HGetAll does not return redis.Nil when no data is found
	if len(values) == 0 {
		return nil, nil
	}
	createdAt, err := time.Parse(time.RFC3339, values["created_at"])
	if err != nil {
		return nil, err
	}
	popularityScore, err := strconv.Atoi(values["popularity_score"])
	if err != nil {
		return nil, err
	}
	return &model.MemeCoin{
		CoinInfo: model.CoinInfo{
			ID:          id,
			Name:        values["name"],
			Description: values["description"],
			CreatedAt:   createdAt,
		},
		CoinScore: model.CoinScore{
			CoinID:          id,
			PopularityScore: popularityScore,
		},
	}, nil
}

func (dao *memeCoinCache) Set(ctx context.Context, memeCoin *model.MemeCoin) error {
	key := dao.buildInfoKey(memeCoin.ID)
	values := map[string]interface{}{
		"name":             memeCoin.Name,
		"description":      memeCoin.Description,
		"popularity_score": memeCoin.PopularityScore,
		"created_at":       memeCoin.CreatedAt.Format(time.RFC3339),
	}
	err := dao.client.HSet(ctx, key, values).Err()
	if err != nil {
		return err
	}
	return dao.client.Expire(ctx, key, time.Minute).Err()
}

func (dao *memeCoinCache) Delete(ctx context.Context, id int64) error {
	key := dao.buildInfoKey(id)
	return dao.client.Del(ctx, key).Err()
}
