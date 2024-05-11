package router

import (
	"net/http"

	"example.com/ginexample/controller"
	"example.com/ginexample/middleware"
	"example.com/ginexample/repositories"
	"example.com/ginexample/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {

	// Authorization
	authorizationRepository := repositories.NewAuthorizationRepositoryImpl(db)
	authorizationService := service.NewAuthorizationServiceImpl(authorizationRepository, validate)
	authorizationController := controller.NewAuthorizationController(authorizationService)

	// Platforms
	platformsRepository := repositories.NewPlatformsRepositoryImpl(db)
	platformsService := service.NewPlatFormsServiceImpl(platformsRepository, validate)
	platformsController := controller.NewPlatformsController(platformsService)

	//attendances
	attendancesService := service.NewAttendancesServiceImpl(validate, db)
	attendanceController := controller.NewAttendancesConroller(attendancesService)

	// Employees
	employeesService := service.NewEmployeesServiceImpl(validate, db)
	employeesController := controller.NewEmployeesController(employeesService)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome to the gin example")
	})
	baseRouter := router.Group("/api")
	{
		baseRouter.POST("/login", authorizationController.LoginHandler)
		platformsRouter := baseRouter.Group("/platform").Use(middleware.Auth())
		{
			platformsRouter.GET("", platformsController.FindAll)
			platformsRouter.GET("/:id", platformsController.FindById)
			platformsRouter.POST("", platformsController.Create)
			platformsRouter.PATCH("/:id", platformsController.Update)
			platformsRouter.DELETE("/:id", platformsController.Delete)
		}
		attendancesRouter := baseRouter.Group("/attendance").Use(middleware.Auth())
		{
			attendancesRouter.GET("", attendanceController.FindAll)
		}

		employeesRouter := baseRouter.Group("/employee").Use(middleware.Auth())
		{
			employeesRouter.GET("", employeesController.FindAllEmployees)
			employeesRouter.GET("/:id", employeesController.FindByIdEmployees)
		}
	}

	return router
}
