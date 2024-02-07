package controller

import (
	"net/http"

	"example.com/ginexample/auth"
	"example.com/ginexample/data/responses"
	"example.com/ginexample/service"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

type AuthorizationController struct {
	authorizationService service.AuthorizationService
}

func NewAuthorizationController(service service.AuthorizationService) *AuthorizationController {
	return &AuthorizationController{
		authorizationService: service,
	}
}

func (controller *AuthorizationController) LoginHandler(ctx *gin.Context) {
	var u responses.AuthorizationResponse

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	authResponse := controller.authorizationService.FindUser("riiamri23", "password123")

	spew.Dump(authResponse)

	var isAuthorization = u.Username == "Check" && u.Password == "123456"
	if isAuthorization {
		tokenString, err := auth.CreateToken(u.Username)

		if err == nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusOK, tokenString)
		} else {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusInternalServerError, "No username found")
		}
	} else {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusUnauthorized, "Username or password wrong")
	}

}
