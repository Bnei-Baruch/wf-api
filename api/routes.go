package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/:root/find", GetRecordsByKV)
	router.GET("/:root/:id", GetRecordByID)
	router.PUT("/:root/:id", PutRecord)
	router.DELETE("/:root/:id", RemoveRecord)
	router.POST("/:root/:id/:key", UpdateRecord)
	router.POST("/:root/:id/:key/:prop", UpdateJsonbRecord)

	router.GET("/trim", GetTrimmed)
	// TODO
	//router.GET("/drim", )
	//router.GET("/drim/:id", )
	//router.GET("/bdika", )
	//router.GET("/jobs_list", )
	//router.GET("/cassette", )
	//router.GET("/states", )
}
