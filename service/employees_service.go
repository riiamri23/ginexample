package service

import (
	"example.com/ginexample/data/data_employee"
	"example.com/ginexample/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type EmployeesService interface {
	FindAll() []data_employee.EmployeeResponse
	FindById(id int) data_employee.EmployeeResponse
	Create(data data_employee.EmployeeResponse)
	Update(data data_employee.EmployeeResponse)
	Delete(id int)
}

type EmployeesServiceImp struct {
	Validate *validator.Validate
	Db       *gorm.DB
}

// Create implements EmployeesService.
func (e *EmployeesServiceImp) Create(data data_employee.EmployeeResponse) {
	panic("unimplemented")
}

// Delete implements EmployeesService.
func (e *EmployeesServiceImp) Delete(id int) {
	panic("unimplemented")
}

// FindAll implements EmployeesService.
func (e *EmployeesServiceImp) FindAll() []data_employee.EmployeeResponse {
	// panic("unimplemented")

	var data []data_employee.Employees

	result := e.Db.Find(&data)

	helpers.ErrorPanic(result.Error)

	var employees []data_employee.EmployeeResponse

	for _, value := range data {
		// date formatting

		tempData := data_employee.EmployeeResponse{
			Id:                           value.Id,
			FirstName:                    &value.Firstname,
			LastName:                     &value.Lastname,
			ContactNo:                    &value.Contact_no,
			OfficialEmail:                &value.Official_email,
			PersonalEmail:                &value.Personal_email,
			IdentityNo:                   &value.Identity_no,
			DateOfBirth:                  helpers.FormatDate(value.Date_of_birth, ""),
			Gender:                       &value.Gender,
			EmergencyContactRelationship: &value.Gender,
			EmergencyContact:             &value.Gender,
			EmerrgencyContactAddress:     &value.Gender,
			City:                         &value.Gender,
			Designation:                  &value.Gender,
			Type:                         &value.Gender,
			Status:                       value.Status,
			EmploymentStatus:             &value.Gender,
			Picture:                      &value.Gender,
			JoiningDate:                  helpers.FormatDate(value.Joining_date, ""),
			ExitDate:                     helpers.FormatDate(value.Exit_date, ""),
			GrossSalary:                  value.Gross_salary,
			Bonus:                        value.Bonus,
			BranchId:                     &value.Gender,
			DepartmentId:                 &value.Gender,
		}

		employees = append(employees, tempData)
	}

	return employees
}

// FindById implements EmployeesService.
func (e *EmployeesServiceImp) FindById(id int) data_employee.EmployeeResponse {
	panic("unimplemented")
}

// Update implements EmployeesService.
func (e *EmployeesServiceImp) Update(data data_employee.EmployeeResponse) {
	panic("unimplemented")
}

func NewEmployeesServiceImpl(validate *validator.Validate, Db *gorm.DB) EmployeesService {
	return &EmployeesServiceImp{
		Validate: validate,
		Db:       Db,
	}
}
