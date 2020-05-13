package main

import (
	"fmt"

	_ "./docs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/patrickmn/go-cache"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Liquidance model
type Liquidance struct {
	Usd float32 `json:"usd" binding:"required"`
	Eur float32 `json:"eur"`
	Btc float32 `json:"btc"`
}

// @title liquides mapping
// @version 1.0
// @description test

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:4421
// @BasePath /api
func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		switch defaultTableName {
		case "liquidances":
			return "public.liquidance"
		default:
			return "public." + defaultTableName
		}

	}

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=Ph@!436623 sslmode=disable")
	if err != nil {
		panic("Can't connect to db")
	}

	defer db.Close()

	host := gin.Default()
	router := host.Group("/api")

	memCache := cache.New(cache.NoExpiration, cache.NoExpiration)

	url := ginSwagger.URL("http://localhost:4421/swagger/doc.json")
	host.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/app/info", getAppInfo)

	router.POST("/liquidity/:provider", authMiddleware, func(context *gin.Context) {
		var data Liquidance
		context.BindJSON(&data)

		db.Create(&data)

		memCache.Set("liquidity", &data, cache.DefaultExpiration)

		context.JSON(200, gin.H{"status": "successfully processed"})
	})

	router.GET("/liquidity", func(context *gin.Context) {
		var liquidances []Liquidance
		var amount int

		query := db.Model(&Liquidance{})

		query.Limit(5).Order("id DESC").Find(&liquidances)
		query.Count(&amount)

		context.JSON(200, gin.H{"data": liquidances, "totalSize": amount})
	})

	router.GET("/liquidity/current", func(context *gin.Context) {
		liquidance, isSuccess := memCache.Get("liquidity")
		if !isSuccess {
			var dbLiq Liquidance
			db.Last(&dbLiq)
			fmt.Println("not from cache")
			context.JSON(200, gin.H{"data": dbLiq})
		} else {
			context.JSON(200, gin.H{"data": liquidance})
		}

	})

	host.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{"message": "resource not found"})
	})

	host.Run(":4421")

	fmt.Println("Server is started on localhost:4421")
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

func authMiddleware(context *gin.Context) {
	auth := context.GetHeader("X-Authorization")
	if auth == "test" {
		context.Next()
	} else {
		context.AbortWithStatusJSON(401, gin.H{"message": "you don't have permission"})
	}
}
