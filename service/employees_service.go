package service

import (
	"example.com/ginexample/data/data_employee"
	"example.com/ginexample/helpers"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type EmployeesService interface {
	FindAll() []data_employee.EmployeeResponse
	FindById(id int) *data_employee.EmployeeResponse
	Create(data data_employee.EmployeeResponse)
	Update(data data_employee.EmployeeResponse)
	Delete(id int)
}

type EmployeesServiceImp struct {
	Validate *validator.Validate
	Db       *gorm.DB
}

func NewEmployeesServiceImpl(validate *validator.Validate, Db *gorm.DB) EmployeesService {
	return &EmployeesServiceImp{
		Validate: validate,
		Db:       Db,
	}
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
func (e *EmployeesServiceImp) FindById(id int) *data_employee.EmployeeResponse {
	// panic("unimplemented")
	var data data_employee.Employees

	result := e.Db.Find(&data, id)
	if result == nil {
		// Handle error if needed
		return nil
	}
	spew.Dump(data.Id)
	if data.Id == 0 {
		return nil
	}

	// helpers.ErrorPanic(result.Error)
	employee := &data_employee.EmployeeResponse{
		Id:                           data.Id,
		FirstName:                    &data.Firstname,
		LastName:                     &data.Lastname,
		ContactNo:                    &data.Contact_no,
		OfficialEmail:                &data.Official_email,
		PersonalEmail:                &data.Personal_email,
		IdentityNo:                   &data.Identity_no,
		DateOfBirth:                  helpers.FormatDate(data.Date_of_birth, ""),
		Gender:                       &data.Gender,
		EmergencyContactRelationship: &data.Gender,
		EmergencyContact:             &data.Gender,
		EmerrgencyContactAddress:     &data.Gender,
		City:                         &data.Gender,
		Designation:                  &data.Gender,
		Type:                         &data.Gender,
		Status:                       data.Status,
		EmploymentStatus:             &data.Gender,
		Picture:                      &data.Gender,
		JoiningDate:                  helpers.FormatDate(data.Joining_date, ""),
		ExitDate:                     helpers.FormatDate(data.Exit_date, ""),
		GrossSalary:                  data.Gross_salary,
		Bonus:                        data.Bonus,
		BranchId:                     &data.Gender,
		DepartmentId:                 &data.Gender,
	}

	return employee
}

// Update implements EmployeesService.
func (e *EmployeesServiceImp) Update(data data_employee.EmployeeResponse) {
	panic("unimplemented")
}
