package controller

import (
	"github.com/c0utin/historinha/data/request"
	"github.com/c0utin/historinha/data/response"
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body     request.CreateUsersRequest true "User request body"
// @Success 200    {object} response.Response
// @Failure 400    {object} response.Response
// @Router /users [post]
func (controller *UsersController) Create(ctx *gin.Context) {
	log.Info().Msg("create user")
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	controller.usersService.Create(createUsersRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   userId path     int                             true "User ID"
// @Param   user   body     request.UpdateUsersRequest true "User request body"
// @Success 200    {object} response.Response
// @Failure 400    {object} response.Response
// @Router /users/{userId} [put]
func (controller *UsersController) Update(ctx *gin.Context) {
	log.Info().Msg("update user")
	updateUsersRequest := request.UpdateUsersRequest{}
	err := ctx.ShouldBindJSON(&updateUsersRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	updateUsersRequest.ID = id

	controller.usersService.Update(updateUsersRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Produce  json
// @Param   userId path     int true "User ID"
// @Success 200    {object} response.Response
// @Failure 400    {object} response.Response
// @Router /users/{userId} [delete]
func (controller *UsersController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete user")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.usersService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindUserById godoc
// @Summary Find a user by ID
// @Description Get a user by ID
// @Tags users
// @Produce  json
// @Param   userId path     int true "User ID"
// @Success 200    {object} response.Response
// @Failure 400    {object} response.Response
// @Router /users/{userId} [get]
func (controller *UsersController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid user")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.usersService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users [get]
func (controller *UsersController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll users")
	userResponse := controller.usersService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

