package app

import (
	"github.com/gin-gonic/gin"
)

//Registers application health endpoints
func RegisterAppHealthEndpoints(router *gin.RouterGroup) {
	router.GET("/app/info", getAppInfo)
}

// @Tags App
// @Summary returns current application's version
// @Description get app version
// @Accept  json
// @Produce  json
// @Success 200
// @Router /app/info [get]
func getAppInfo(context *gin.Context) {
	context.JSON(200, gin.H{"message": "version 1.0"})
}
