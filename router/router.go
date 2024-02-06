package router

import (
	"net/http"

	"example.com/ginexample/controller"
	"example.com/ginexample/middleware"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(platformsController *controller.PlatformsController) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	{
		baseRouter.POST("/login", controller.LoginHandler)
		platformsRouter := baseRouter.Group("/platform").Use(middleware.Auth())
		{
			platformsRouter.GET("", platformsController.FindAll)
			platformsRouter.GET("/:id", platformsController.FindById)
			platformsRouter.POST("", platformsController.Create)
			platformsRouter.PATCH("/:id", platformsController.Update)
			platformsRouter.DELETE("/:id", platformsController.Delete)
		}
	}

	return router
}
