package controller

import (
	"net/http"
	"strconv"

	"example.com/ginexample/data/requests"
	"example.com/ginexample/data/responses"
	"example.com/ginexample/helpers"
	"example.com/ginexample/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PlatformsController struct {
	platformsService service.PlatformsService
}

func NewPlatformsController(service service.PlatformsService) *PlatformsController {
	return &PlatformsController{
		platformsService: service,
	}
}

// CreatePlatforms		godoc
// @Summary			Create Platforms
// @Description		Save Platforms data in Db.
// @Param			Platforms body request.CreatePlatformsRequest true "Create platform"
// @Produce			application/json
// @Platforms		Platforms
// @Success			200 {object} response.Response{}
// @Router			/platform [post]
func (controller *PlatformsController) Create(ctx *gin.Context) {
	createPlatformsRequest := requests.CreatePlatformsRequest{}
	// spew.Dump(ctx)
	err := ctx.ShouldBindJSON(&createPlatformsRequest)
	helpers.ErrorPanic(err)

	controller.platformsService.Create(createPlatformsRequest)
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PlatformsController) Update(ctx *gin.Context) {
	log.Info().Msg("update platforms")
	updatePlatformsRequest := requests.UpdatePlatformsRequest{}
	err := ctx.ShouldBindJSON(&updatePlatformsRequest)
	helpers.ErrorPanic(err)

	platformId := ctx.Param("id")
	id, err := strconv.Atoi(platformId)
	helpers.ErrorPanic(err)
	updatePlatformsRequest.Id = id

	controller.platformsService.Update(updatePlatformsRequest)

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PlatformsController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete platforms")
	platformId := ctx.Param("id")
	id, err := strconv.Atoi(platformId)
	helpers.ErrorPanic(err)
	controller.platformsService.Delete(id)

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PlatformsController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid platforms")
	platformId := ctx.Param("id")
	id, err := strconv.Atoi(platformId)
	helpers.ErrorPanic(err)

	tagResponse := controller.platformsService.FindById(id)

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// CreatePlatforms		godoc
// @Summary			Get All Platforms
// @Description		Save Platforms data in Db.
// @Param			Platforms body empty
// @Produce			application/json
// @Platforms		Platforms
// @Success			200 {object} response.Response{}
// @Router			/platform
func (controller *PlatformsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll platforms")
	tagResponse := controller.platformsService.FindAll()
	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
