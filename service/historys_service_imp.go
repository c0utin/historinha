package service

import (
	"github.com/c0utin/historinha/data/request"
	"github.com/c0utin/historinha/data/response"
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/model"
	"github.com/c0utin/historinha/repository"

	"github.com/go-playground/validator/v10"
)

type HistorysServiceImpl struct {
	HistorysRepository repository.HistorysRepository
	Validate           *validator.Validate
}

func NewHistorysServiceImpl(historysRepository repository.HistorysRepository, validate *validator.Validate) HistorysService {
	return &HistorysServiceImpl{
		HistorysRepository: historysRepository,
		Validate:            validate,
	}
}

// Create implements HistorysService
func (h *HistorysServiceImpl) Create(history request.CreateHistoryRequest) {
	err := h.Validate.Struct(history)
	helper.ErrorPanic(err)

	historyModel := model.Historys{
		Name:   history.Name,
		History: history.History,
		UserID: history.UserID,
	}
	h.HistorysRepository.Save(historyModel)
}

// Delete implements HistorysService
func (h *HistorysServiceImpl) Delete(historyId int) {
	h.HistorysRepository.Delete(historyId)
}

// FindAll implements HistorysService
func (h *HistorysServiceImpl) FindAll() []response.HistoryResponse {
	result := h.HistorysRepository.FindAll()

	var histories []response.HistoryResponse
	for _, value := range result {
		history := response.HistoryResponse{
			ID:      value.ID,
			Name:    value.Name,
			History: value.History,
			UserID:  value.UserID,
		}
		histories = append(histories, history)
	}

	return histories
}

// FindById implements HistorysService
func (h *HistorysServiceImpl) FindById(historyId int) response.HistoryResponse {
    historyData, err := h.HistorysRepository.FindById(historyId)
    helper.ErrorPanic(err)

    historyResponse := response.HistoryResponse{
        ID:      historyData.ID,
        Name:    historyData.Name,
        History: historyData.History,
        UserID:  historyData.UserID,
    }
    return historyResponse
}



// Update implements HistorysService
func (h *HistorysServiceImpl) Update(history request.UpdateHistoryRequest) {
	historyData, err := h.HistorysRepository.FindById(history.ID)
	helper.ErrorPanic(err)

	historyData.Name = history.Name
	historyData.History = history.History

	h.HistorysRepository.Update(historyData)
}

