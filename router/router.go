package router

import (
	"net/http"

	"example.com/ginexample/controller"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(platformsController *controller.PlatformsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	platformsRouter := baseRouter.Group("/platform")
	platformsRouter.GET("", platformsController.FindAll)
	platformsRouter.GET("/:id", platformsController.FindById)
	platformsRouter.POST("", platformsController.Create)
	platformsRouter.PATCH("/:id", platformsController.Update)
	platformsRouter.DELETE("/:id", platformsController.Delete)

	return router
}
