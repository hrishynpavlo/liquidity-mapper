package liquidity

import (
	"../common"
	"../middlewares"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

//Registers liquidity endpoints
func RegisterLiquidityEndpoints(router *gin.RouterGroup) {
	router.POST("/liquidity/:provider", middlewares.Auth, Create)
	router.GET("/liquidity", GetAll)
	router.GET("/liquidity/latest", GetLatest)
}

// @Tags Liquidity
// @Summary creates new liquidity
// @Description saves in database new record and sets it in cache
// @Accept  json
// @Produce  json
// @Success 200
// @Security ApiKeyAuth
// @Router /liquidity/:provider [post]
// @Param :provider path string true "provider"
// @Param liquidity body Liquidance true "liquidity"
func Create(context *gin.Context) {
	var data Liquidance
	error := context.BindJSON(&data)
	if error != nil {
		context.AbortWithStatus(400)
		return
	}

	common.GetDbInstance().Create(&data)

	common.GetCacheInstance().Set("liquidity", &data, cache.DefaultExpiration)

	context.JSON(200, gin.H{"status": "successfully processed"})
}

// @Tags Liquidity
// @Summary returns 5 latest liquidities
// @Description returns 5 latest liquidities
// @Accept  json
// @Produce  json
// @Success 200
// @Router /liquidity [get]
func GetAll(context *gin.Context) {
	var liquidances []Liquidance
	var amount int

	query := common.GetDbInstance().Model(&Liquidance{})

	query.Limit(5).Order("id DESC").Find(&liquidances)
	query.Count(&amount)

	context.JSON(200, gin.H{"data": liquidances, "totalSize": amount})
}

// @Tags Liquidity
// @Summary returns latest liquidities
// @Description returns 5 latest liquidities
// @Accept  json
// @Produce  json
// @Success 200
// @Router /liquidity/latest [get]
func GetLatest(context *gin.Context) {
	memCache := common.GetCacheInstance()
	liquidance, isSuccess := memCache.Get("liquidity")
	if !isSuccess {
		var dbLiq Liquidance
		common.GetDbInstance().Last(&dbLiq)
		memCache.Set("liquidity", &dbLiq, cache.DefaultExpiration)
		context.JSON(200, gin.H{"data": dbLiq})
	} else {
		context.JSON(200, gin.H{"data": liquidance})
	}
}
