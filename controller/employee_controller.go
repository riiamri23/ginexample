package controller

import (
	"net/http"
	"strconv"

	"example.com/ginexample/data/data_employee"
	"example.com/ginexample/data/responses"
	"example.com/ginexample/helpers"
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

func (controller *EmployeesController) CreateEmployees(ctx *gin.Context) {
	createEmployeeRequest := data_employee.Employees{}

	err := ctx.ShouldBindJSON(&createEmployeeRequest)
	helpers.ErrorPanic(err)

	controller.employeesService.Create(createEmployeeRequest)

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Successfully save the employee data",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
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

func (controller *EmployeesController) FindByIdEmployees(ctx *gin.Context) {

	log.Info().Msg("Find By Id Employee")
	employeeId := ctx.Param("id")
	id, err := strconv.Atoi(employeeId)
	helpers.ErrorPanic(err)

	tagResponse := controller.employeesService.FindById(id)
	if tagResponse != nil {

		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   tagResponse,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)
	} else {

		webResponse := responses.Response{
			Code:   http.StatusNoContent,
			Status: "Employee is not found",
			Data:   tagResponse,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)
	}
}
