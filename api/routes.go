package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/:root/find", GetRecordsByKV)
	router.GET("/:root/:id", GetRecordByID)
	router.PUT("/:root/:id", PutRecord)
	router.DELETE("/:root/:id", RemoveRecord)
	router.POST("/:root/:id/:key", PostRecordJSON)
	router.POST("/:root/:id/:key/:st", UpdateRecordState)

	router.GET("/trim", GetTrimmed)
}
