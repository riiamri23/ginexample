package controller

import (
	"net/http"

	"example.com/ginexample/data/responses"
	"example.com/ginexample/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AttendancesController struct {
	attendancesService service.AttendancesService
}

func NewAttendancesConroller(service service.AttendancesService) *AttendancesController {
	return &AttendancesController{attendancesService: service}
}

func Create(ctx *gin.Context) {
}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}

func FindById(ctx *gin.Context) {

}

func (controller *AttendancesController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll attendance")
	tagResponse := controller.attendancesService.FindAll()

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
