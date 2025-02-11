package model

import (
	"time"

	"gorm.io/gorm"
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
	ID          int64     `json:"-" gorm:"primaryKey;column:id;autoIncrement:false"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null;column:name;uniqueIndex"`
	Description string    `json:"description" gorm:"type:text;column:description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at"`
}

func (*CoinInfo) TableName() string {
	return TableCoinInfo.String()
}

type CoinScore struct {
	CoinID          int64 `json:"-" gorm:"column:coin_id;not null;uniqueIndex:idx_coin_id"`
	PopularityScore int   `json:"popularity_score" gorm:"type:int;not null;column:popularity_score"`
}

func (*CoinScore) TableName() string {
	return TableCoinScore.String()
}

func MemeCoinMigrate(db *gorm.DB) error {
	if db.Migrator().HasTable(&CoinInfo{}) && db.Migrator().HasTable(&CoinScore{}) {
		return nil
	}
	return db.AutoMigrate(&CoinInfo{}, &CoinScore{})
}
