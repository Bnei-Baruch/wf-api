package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	state := router.Group("state")
	state.GET("/:tag", GetStateByTag)
	state.GET("/:tag/:id", GetStateByID)
	state.GET("/:tag/:id/:prop", GetStateByProp)
	state.DELETE("/:tag/:id", RemoveStateByID)
	state.PUT("/:tag/:id", PutStateByID)
	state.POST("/:tag/:id", PutStateByID)
	state.PUT("/:tag/:id/:prop", PutStateByProp)
	state.DELETE("/:tag/:id/:prop", RemoveStateByProp)
	router.GET("/states", GetStates)

	router.GET("/source/file_uid/:uid", GetSourceByUID)

	router.GET("/:root/find", V1GetRecordsByKV)
	router.GET("/:root/kv", V2GetRecordsByKV)
	router.GET("/:root/js/:prop", GetRecordsByJSON)
	router.GET("/:root/:id", GetRecordByID)
	router.PUT("/:root/:id", PutRecord)
	router.DELETE("/:root/:id", RemoveRecord)
	router.POST("/:root/:id/:key", UpdateRecord)
	router.POST("/:root/:id/:key/:prop", UpdateJsonbRecord)

	router.GET("/trim", GetTrimmed)
	router.GET("/jobs", GetJobs)
	router.GET("/drim", GetDgima)
	router.GET("/bdika", GetAricha)
	// TODO
	//router.GET("/drim/:id", )
	//router.GET("/jobs_list", )
	//router.GET("/cassette", )
}
