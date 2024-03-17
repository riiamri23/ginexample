package middleware

import (
	"example.com/ginexample/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		// spew.Dump(tokenString)
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := auth.VerifyToken(tokenString)

		if err != nil {
			context.JSON(401, gin.H{"error token": err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}
