package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/state/:id", GetStateByID)
	router.GET("/state/:id/:prop", GetStateByProp)
	router.DELETE("/state/:id", RemoveStateProp)

	router.GET("/:root/find", V1GetRecordsByKV)
	router.GET("/:root/kv", V2GetRecordsByKV)
	router.GET("/:root/js/:prop", GetRecordsByJSON)
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
