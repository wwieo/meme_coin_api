package dao

import "meme_coin_api/service/internal/model"

type MemeCoinDao interface {
	Create(coinInfo *model.CoinInfo, coinScore *model.CoinScore) error
	Get(id int64) (*model.MemeCoin, error)
	UpdateDescription(id int64, description string) error
	IncreasePopularityScore(id int64) error
	DeleteMemeCoin(id int64) error
}
