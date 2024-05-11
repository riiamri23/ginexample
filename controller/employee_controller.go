package controller

import (
	"net/http"

	"example.com/ginexample/data/responses"
	"example.com/ginexample/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type EmployeesController struct {
	employeesService service.EmployeesService
}

func NewEmployeesController(service service.EmployeesService) *EmployeesController {
	return &EmployeesController{employeesService: service}
}

func CreateEmployees(ctx *gin.Context) {

}

func UpdateEmployees(ctx *gin.Context) {

}

func DeleteEmployees(ctx *gin.Context) {

}

func (controller *EmployeesController) FindAllEmployees(ctx *gin.Context) {
	log.Info().Msg("FindAll Employee")
	tagResponse := controller.employeesService.FindAll()

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
