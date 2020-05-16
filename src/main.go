package main

import (
	"fmt"

	"./common"
	"./configuration"
	_ "./docs"
	app "./health"
	"./liquidity"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title fluent liquidity mapper v1.0
// @host localhost:4421
// @BasePath /api
// @contact.name Pavlo Hrishyn
// @contact.email pashagrishyn@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Authorization
func main() {

	configuration.SetUp()
	common.DbInit()
	common.CacheInit()

	host := gin.Default()
	host.Use(cors.Default())
	router := host.Group("/api")

	host.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{"message": "resource not found"})
	})

	url := ginSwagger.URL("http://localhost:4421/swagger/doc.json")
	host.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	app.RegisterAppHealthEndpoints(router)
	liquidity.RegisterLiquidityEndpoints(router)

	host.Run(":4421")

	fmt.Println("Server is started on localhost:4421")
}
