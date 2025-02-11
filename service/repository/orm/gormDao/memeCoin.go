package gormDao

import (
	"errors"
	"gorm.io/gorm"
	"meme_coin_api/service/internal/errorx"
	"meme_coin_api/service/internal/model"
	"meme_coin_api/service/repository/orm"
)

type memeCoinDao struct {
	db *gorm.DB
}

func NewMemeCoinDao(db *gorm.DB) orm.MemeCoinDao {
	return &memeCoinDao{
		db: db,
	}
}

func (dao *memeCoinDao) Create(coinInfo *model.CoinInfo) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(model.TableCoinInfo.String()).
			Create(coinInfo).Error; err != nil {
			if errorx.IsMySQLDuplicateEntry(err) {
				return errorx.RecordExisted
			}
			return err
		}

		if err := tx.Table(model.TableCoinScore.String()).
			Create(&model.CoinScore{CoinID: coinInfo.ID}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (dao *memeCoinDao) Get(id int64) (*model.MemeCoin, error) {
	var memeCoin *model.MemeCoin
	if err := dao.db.Table(model.TableCoinInfo.String()).
		Select("coin_info.id, coin_info.name, coin_info.description, coin_info.created_at, coin_score.popularity_score").
		Joins("join coin_score on coin_info.id = coin_score.coin_id").
		Where("coin_info.id = ?", id).
		First(&memeCoin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.NoRecord
		}
		return nil, err
	}
	return memeCoin, nil
}

func (dao *memeCoinDao) UpdateDescription(id int64, description string) error {
	if err := dao.db.Table(model.TableCoinInfo.String()).
		Model(&model.CoinInfo{}).
		Where("id = ?", id).
		Update("description", description).
		Error; err != nil {
		return err
	}
	return nil
}

func (dao *memeCoinDao) IncreasePopularityScore(id int64) error {
	if err := dao.db.Table(model.TableCoinScore.String()).
		Model(&model.CoinScore{}).
		Where("coin_id = ?", id).
		Update("popularity_score", gorm.Expr("popularity_score + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (dao *memeCoinDao) DeleteMemeCoin(id int64) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(model.TableCoinScore.String()).
			Delete(&model.CoinScore{}, "coin_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Table(model.TableCoinInfo.String()).
			Delete(&model.CoinInfo{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}
