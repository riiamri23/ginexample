package controller

import (
	"fmt"
	"net/http"

	"example.com/ginexample/auth"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(ctx *gin.Context) {
	var u User

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	fmt.Printf("The user request value %v", u)

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

	return
}
