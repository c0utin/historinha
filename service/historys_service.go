package service

import (
	"github.com/c0utin/historinha/data/request"
	"github.com/c0utin/historinha/data/response"
)

type HistorysService interface {
	Create(history request.CreateHistoryRequest)
	Update(history request.UpdateHistoryRequest)
	Delete(historyId int)
	FindById(historyId int) response.HistoryResponse
	FindAll() []response.HistoryResponse
}

