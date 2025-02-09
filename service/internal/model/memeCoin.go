package model

import (
	"time"
)

const (
	TableCoinInfo  TableName = "coin_info"
	TableCoinScore TableName = "coin_score"
)

type MemeCoin struct {
	CoinInfo
	CoinScore
}

type CoinInfo struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CoinScore struct {
	CoinID          int64 `json:"coin_id"`
	PopularityScore int   `json:"popularity_score"`
}
