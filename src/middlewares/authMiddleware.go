package middlewares

import (
	"github.com/gin-gonic/gin"
)

//Auth returns
func Auth(context *gin.Context) {
	auth := context.GetHeader("X-Authorization")
	if auth == "test" {
		context.Next()
	} else {
		context.AbortWithStatusJSON(401, gin.H{"message": "you don't have permission"})
	}
}
