package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/get_users", GetUsers)
	router.GET("/get_user/:id", GetUser)
	router.POST("/create_user", CreateUser)

	router.GET("/ingest/find", GetIngestByKV)
	router.GET("/ingest/:id", GetIngestByID)
	router.PUT("/ingest/:id", PutIngest)
	router.POST("/ingest/:id/wfstatus/:key", UpdateIngestState)
	router.POST("/ingest/:id/:key", PostIngestJSON)
	router.DELETE("/ingest/:id", RemoveIngest)

	router.GET("/trimmer/find", GetTrimmerByKV)
	router.GET("/trimmer/:id", GetTrimmerByID)
	router.PUT("/trimmer/:id", PutTrimmer)
	router.POST("/trimmer/:id/wfstatus/:key", UpdateTrimmerState)
	router.POST("/trimmer/:id/:key", PostTrimmerJSON)
	router.DELETE("/trimmer/:id", RemoveTrimmer)
	router.GET("/trim", GetTrimmed)

	router.GET("/kmedia/find", GetKmediaByKV)
	router.GET("/kmedia/:id", GetKmediaByID)
	router.PUT("/kmedia/:id", PutKmedia)
	router.DELETE("/kmedia/:id", RemoveKmedia)

	router.GET("/products/find", GetProductsByDF)
}
