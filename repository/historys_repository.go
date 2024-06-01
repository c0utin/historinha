package repository 

import "github.com/c0utin/historinha/model"

type HistorysRepository interface {
	Save(historys model.Historys)
	Update(historys model.Historys)
	Delete(historyId int)
	FindById(historyId int) (historys model.Historys, err error)
	FindAll() []model.Historys
}
