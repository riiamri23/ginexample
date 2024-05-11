package service

import (
	"example.com/ginexample/data/data_attendance"
	"example.com/ginexample/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AttendancesService interface {
	FindAll() []data_attendance.AttendanceResponse
	FindById(id int) data_attendance.AttendanceResponse
	Create(attendances data_attendance.CreateAttendanceRequest)
	Update(attendances data_attendance.UpdateAttendanceRequest)
	Delete(id int)
}

type AttendanceServiceImp struct {
	Validate *validator.Validate
	Db       *gorm.DB
}

func NewAttendancesServiceImpl(validate *validator.Validate, Db *gorm.DB) AttendancesService {
	return &AttendanceServiceImp{
		Validate: validate,
		Db:       Db,
	}
}

// Create implements AttendancesService.
func (a *AttendanceServiceImp) Create(attendances data_attendance.CreateAttendanceRequest) {
	panic("unimplemented")
}

// Delete implements AttendancesService.
func (a *AttendanceServiceImp) Delete(id int) {
	panic("unimplemented")
}

// FindAll implements AttendancesService.
func (a *AttendanceServiceImp) FindAll() []data_attendance.AttendanceResponse {
	var data []data_attendance.Attendances

	result := a.Db.Find(&data)

	helpers.ErrorPanic(result.Error)

	var attendances []data_attendance.AttendanceResponse
	for _, value := range data {
		tempData := data_attendance.AttendanceResponse{
			Id:         value.Id,
			EmployeeId: value.Employee_Id,
		}

		attendances = append(attendances, tempData)
	}

	return attendances
}

// FindById implements AttendancesService.
func (a *AttendanceServiceImp) FindById(id int) data_attendance.AttendanceResponse {
	panic("unimplemented")
}

// Update implements AttendancesService.
func (a *AttendanceServiceImp) Update(attendances data_attendance.UpdateAttendanceRequest) {
	panic("unimplemented")
}
