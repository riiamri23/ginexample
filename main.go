package main

import (
	"context"
	"net/http"
	"syscall"

	"example.com/ginexample/config"
	"example.com/ginexample/controller"
	"example.com/ginexample/helpers"
	"example.com/ginexample/model"
	"example.com/ginexample/repositories"
	"example.com/ginexample/router"
	"example.com/ginexample/service"

	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")
	// Database config
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("platforms").AutoMigrate(&model.Platforms{})

	// Authorization
	authorizationRepository := repositories.NewAuthorizationRepositoryImpl(db)
	authorizationService := service.NewAuthorizationServiceImpl(authorizationRepository, validate)
	authorizationController := controller.NewAuthorizationController(authorizationService)

	// Platforms
	platformsRepository := repositories.NewPlatformsRepositoryImpl(db)
	platformsService := service.NewPlatFormsServiceImpl(platformsRepository, validate)
	platformsController := controller.NewPlatformsController(platformsService)

	// Router
	routes := router.InitRouter(authorizationController, platformsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helpers.ErrorPanic(err)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	log.Printf("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown:", err)
	}

	log.Printf("Server Stopped")
}
