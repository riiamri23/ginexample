package main

import (
	"net/http"

	"example.com/ginexample/config"
	"example.com/ginexample/controller"
	"example.com/ginexample/helpers"
	"example.com/ginexample/model"
	"example.com/ginexample/repositories"
	"example.com/ginexample/router"
	"example.com/ginexample/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")
	// Database config
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("platforms").AutoMigrate(&model.Platforms{})

	// Repository
	platformsRepository := repositories.NewPlatformsRepositoryImpl(db)

	// Service
	platformsService := service.NewPlatFormsServiceImpl(platformsRepository, validate)

	// Controller
	platformsController := controller.NewPlatformsController(platformsService)

	// routes := gin.Default()

	// routes.GET("", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "Welcome Home")
	// })

	// Router
	routes := router.NewRouter(platformsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	// helper.ErrorPanic(err)
	helpers.ErrorPanic(err)
}
