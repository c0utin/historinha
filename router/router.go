package router

import (
	"github.com/c0utin/historinha/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewUsersRouter(usersController *controller.UsersController) *gin.Engine {
	router := gin.Default()

	// swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	usersRouter := baseRouter.Group("/users")
	usersRouter.GET("", usersController.FindAll)
	usersRouter.GET("/:userId", usersController.FindById)
	usersRouter.POST("", usersController.Create)
	usersRouter.PATCH("/:userId", usersController.Update)
	usersRouter.DELETE("/:userId", usersController.Delete)

	return router
}
