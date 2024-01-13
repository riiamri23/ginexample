package service

import (
	"example.com/ginexample/data/requests"
	"example.com/ginexample/data/responses"
	"example.com/ginexample/helpers"
	"example.com/ginexample/model"
	"example.com/ginexample/repositories"

	"github.com/go-playground/validator/v10"
)

type PlatformsService interface {
	FindAll() []responses.PlatformsResponse
	FindById(id int) responses.PlatformsResponse
	Create(platforms requests.CreatePlatformsRequest)
	Update(platforms requests.UpdatePlatformsRequest)
	Delete(id int)
}

type PlatformsServiceImp struct {
	PlatformsRepository repositories.PlatformsRepository
	Validate            *validator.Validate
}

func NewPlatFormsServiceImpl(platformsRepository repositories.PlatformsRepository, validate *validator.Validate) PlatformsService {
	return &PlatformsServiceImp{
		PlatformsRepository: platformsRepository,
		Validate:            validate,
	}
}

// Add implements PlatformsService
func (t *PlatformsServiceImp) Create(platforms requests.CreatePlatformsRequest) {
	// panic("unimplemented")
	err := t.Validate.Struct(platforms)
	helpers.ErrorPanic(err)
	platformsModel := model.Platforms{
		Name: platforms.Name,
	}
	t.PlatformsRepository.Create(platformsModel)
}

// Delete implements PlatformsService
func (t *PlatformsServiceImp) Delete(id int) {
	t.PlatformsRepository.Delete(id)
}

// FindAll implements PlatformsService
func (t *PlatformsServiceImp) FindAll() []responses.PlatformsResponse {
	result := t.PlatformsRepository.FindAll()

	var platforms []responses.PlatformsResponse
	for _, value := range result {
		platform := responses.PlatformsResponse{
			Id:   value.Id,
			Name: value.Name,
		}

		platforms = append(platforms, platform)
	}

	return platforms
}

// FindById implements PlatformsService
func (t *PlatformsServiceImp) FindById(id int) responses.PlatformsResponse {
	platformData, err := t.PlatformsRepository.FindById(id)
	helpers.ErrorPanic(err)

	platformResponse := responses.PlatformsResponse{
		Id:   platformData.Id,
		Name: platformData.Name,
	}

	return platformResponse
}

// Update implements PlatformsService
func (t *PlatformsServiceImp) Update(platforms requests.UpdatePlatformsRequest) {
	platformData, err := t.PlatformsRepository.FindById(platforms.Id)
	helpers.ErrorPanic(err)

	platformData.Name = platforms.Name
	t.PlatformsRepository.Update(platformData)

}
