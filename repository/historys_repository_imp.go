package repository

import (
	"errors"
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/model"
	"gorm.io/gorm"
)

type HistorysRepositoryImpl struct {
	Db *gorm.DB
}

func NewHistorysRepositoryImpl(Db *gorm.DB) HistorysRepository {
	return &HistorysRepositoryImpl{Db: Db}
}

// Save implements HistorysRepository
func (r *HistorysRepositoryImpl) Save(history model.Historys) {
	result := r.Db.Create(&history)
	helper.ErrorPanic(result.Error)
}

// Update implements HistorysRepository
func (r *HistorysRepositoryImpl) Update(history model.Historys) {
	result := r.Db.Model(&history).Updates(history)
	helper.ErrorPanic(result.Error)
}

// Delete implements HistorysRepository
func (r *HistorysRepositoryImpl) Delete(historyId int) {
	var history model.Historys
	result := r.Db.Where("id = ?", historyId).Delete(&history)
	helper.ErrorPanic(result.Error)
}

// FindById implements HistorysRepository
func (r *HistorysRepositoryImpl) FindById(historyId int) (history model.Historys, err error) {
	var h model.Historys
	result := r.Db.First(&h, historyId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return h, errors.New("history not found")
		}
		helper.ErrorPanic(result.Error)
	}
	return h, nil
}

// FindAll implements HistorysRepository
func (r *HistorysRepositoryImpl) FindAll() []model.Historys {
	var histories []model.Historys
	result := r.Db.Find(&histories)
	helper.ErrorPanic(result.Error)
	return histories
}

