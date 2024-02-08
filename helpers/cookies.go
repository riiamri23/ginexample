package helpers

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

func SetCookieHandler(c *gin.Context, key string, value string) {
	c.SetCookie(key, value, 3600, "/", "localhost", false, true)
}
