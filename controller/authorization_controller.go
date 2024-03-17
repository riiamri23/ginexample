package controller

import (
	"net/http"

	"example.com/ginexample/auth"
	"example.com/ginexample/data/responses"
	"example.com/ginexample/helpers"

	// "example.com/ginexample/model"
	"example.com/ginexample/service"
	// "github.com/davecgh/go-spew/spew"
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

	authResponse := controller.authorizationService.FindUser(u.Username, u.Password)

	// spew.Dump(authResponse.Username)

	if authResponse.Username != "" {
		tokenString, err := auth.CreateToken(u.Username)
		authResponse.Jwtkey = tokenString

		if err == nil {
			helpers.SetCookieHandler(ctx, "Username", u.Username)
			helpers.SetCookieHandler(ctx, "AuthorizationKey", tokenString)

			// spew.Dump(ctx.Cookie("Username"))
			// spew.Dump(ctx.Cookie("AuthorizationKey"))

			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusOK, authResponse)
		} else {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusInternalServerError, "No username found")
		}
	} else {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusUnauthorized, "Username or password wrong")
	}

}
