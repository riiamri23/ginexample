package service

import (
	"example.com/ginexample/data/responses"
	"example.com/ginexample/helpers"
	"example.com/ginexample/repositories"
	"github.com/go-playground/validator/v10"
)

type AuthorizationService interface {
	FindUser(username string, password string) responses.AuthorizationResponse
}

type AuthorizationServiceImp struct {
	AuthorizationRepository repositories.AuthorizationRepository
	Validate                *validator.Validate
}

func NewAuthorizationServiceImpl(authorizationRepository repositories.AuthorizationRepository, validate *validator.Validate) AuthorizationService {
	return &AuthorizationServiceImp{
		AuthorizationRepository: authorizationRepository,
		Validate:                validate,
	}
}

func (t *AuthorizationServiceImp) FindUser(username string, password string) responses.AuthorizationResponse {
	result, err := t.AuthorizationRepository.FindUser(username, password)

	helpers.ErrorPanic(err)

	userResponse := responses.AuthorizationResponse{
		Username: result.Username,
		Password: result.Password,
	}

	return userResponse

}
