package caches

import (
	"context"
	"meme_coin_api/service/internal/model"
)

type MemeCoinCache interface {
	Exist(ctx context.Context, id int64) (bool, error)
	Get(ctx context.Context, id int64) (*model.MemeCoin, error)
	Set(ctx context.Context, coin *model.MemeCoin) error
	Delete(ctx context.Context, id int64) error
}
