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

	//router.GET("/ingest/find", GetIngestByKV)
	//router.GET("/ingest/:id", GetIngestByID)
	//router.PUT("/ingest/:id", PutIngest)
	//router.POST("/ingest/:id/:st/:key", UpdateIngestState)
	//router.POST("/ingest/:id/:key", PostIngestJSON)
	//router.DELETE("/ingest/:id", RemoveIngest)
	//
	//router.GET("/trimmer/find", GetTrimmerByKV)
	//router.GET("/trimmer/:id", GetTrimmerByID)
	//router.PUT("/trimmer/:id", PutTrimmer)
	//router.POST("/trimmer/:id/:st/:key", UpdateTrimmerState)
	//router.POST("/trimmer/:id/:key", PostTrimmerJSON)
	//router.DELETE("/trimmer/:id", RemoveTrimmer)
	//router.GET("/trim", GetTrimmed)
	//
	//router.GET("/products/find", GetProductsByDF)
	//router.GET("/products/:id", GetProductByID)
	//router.PUT("/products/:id", PutProduct)
	//router.POST("/products/:id/:st/:key", UpdateProductState)
	//router.POST("/products/:id/:key", PostProductJSON)
	//router.DELETE("/products/:id", RemoveProduct)
	//
	//router.GET("/kmedia/find", GetKmediaByKV)
	//router.GET("/kmedia/:id", GetKmediaByID)
	//router.PUT("/kmedia/:id", PutKmedia)
	//router.DELETE("/kmedia/:id", RemoveKmedia)
}
