package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/c0utin/historinha/data/request"
	"github.com/c0utin/historinha/data/response"
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/model"
	"github.com/c0utin/historinha/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HistorysController struct {
	historysService service.HistorysService
}

func NewHistorysController(service service.HistorysService) *HistorysController {
	return &HistorysController{
		historysService: service,
	}
}

func (controller *HistorysController) Create(ctx *gin.Context) {
	log.Info().Msg("create history")
	createHistoryRequest := request.CreateHistoryRequest{}
	err := ctx.ShouldBindJSON(&createHistoryRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	// Call Gemini API to generate the history content
	historyContent, err := callGeminiAPI(createHistoryRequest.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
		return
	}

	history := model.Historys{
		Name:    createHistoryRequest.Name,
		History: historyContent,
		UserID:  createHistoryRequest.UserID,
	}

	controller.historysService.Create(history)

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   history,
	})
}

func (controller *HistorysController) Update(ctx *gin.Context) {
	log.Info().Msg("update history")
	updateHistoryRequest := request.UpdateHistoryRequest{}
	err := ctx.ShouldBindJSON(&updateHistoryRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	historyId, err := strconv.Atoi(ctx.Param("historyId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	history := model.Historys{
		ID:     historyId,
		Name:   updateHistoryRequest.Name,
		UserID: updateHistoryRequest.UserID,
	}

	controller.historysService.Update(history)

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	})
}

func (controller *HistorysController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete history")
	historyId, err := strconv.Atoi(ctx.Param("historyId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	controller.historysService.Delete(historyId)

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	})
}

func (controller *HistorysController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid history")
	historyId, err := strconv.Atoi(ctx.Param("historyId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	historyResponse := controller.historysService.FindById(historyId)

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   historyResponse,
	})
}

func (controller *HistorysController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll histories")
	historyResponse := controller.historysService.FindAll()

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   historyResponse,
	})
}

func callGeminiAPI(name string) (string, error) {
	apiURL := os.Getenv("GEMINI_API_URL")
	return fmt.Sprintf("Generated history for %s using %s", name, apiURL), nil
}
