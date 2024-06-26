package main

import (
	"net/http"

	"github.com/c0utin/historinha/config"
	"github.com/c0utin/historinha/controller"
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/model"
	"github.com/c0utin/historinha/repository"
	"github.com/c0utin/historinha/service"
	"github.com/c0utin/historinha/router"

	_ "github.com/c0utin/historinha/docs"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title Historinha
// @version 1.0
// @description historinhas muito loko
// @host localhost:42069
// @BasePath /api
func main() {
	log.Info().Msg("Server running!")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})
//	db.Table("historys").AutoMigrate(&model.Historys{}) 

	usersRepository := repository.NewUsersRepositoryImpl(db)
//	historysRepository := repository.NewHistorysRepositoryImpl(db) 

	usersService := service.NewUsersServiceImpl(usersRepository, validate)
//	historysService := service.NewHistorysServiceImpl(historysRepository, validate) 

	usersController := controller.NewUsersController(usersService)
//	historysController := controller.NewHistorysController(historysService)

	routes := router.NewRouter(usersController)
	server := &http.Server{
		Addr:    ":42069",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

