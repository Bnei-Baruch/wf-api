package api

import (
	"github.com/Bnei-Baruch/wf-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var list = map[string]interface{}{
	"ingest":   []models.Ingest{},
	"trimmer":  []models.Trimmer{},
	"products": []models.Product{},
	"state":    []models.State{},
	"kmedia":   []models.Kmedia{},
	"archive":  []models.Archive{},
	"capture":  []models.Capture{},
	"convert":  []models.Convert{},
	"carbon":   []models.Carbon{},
	"insert":   []models.Insert{},
	"source":   []models.Source{},
	"dgima":    []models.Dgima{},
	"aricha":   []models.Aricha{},
	"jobs":     []models.Job{},
	"files":    []models.File{},
	"cloud":    []models.Cloud{},
	"users":    []models.User{},
	"label":    []models.Label{},
}

var recd = map[string]interface{}{
	"ingest":   &models.Ingest{},
	"trimmer":  &models.Trimmer{},
	"products": &models.Product{},
	"state":    &models.State{},
	"kmedia":   &models.Kmedia{},
	"archive":  &models.Archive{},
	"capture":  &models.Capture{},
	"convert":  &models.Convert{},
	"carbon":   &models.Carbon{},
	"insert":   &models.Insert{},
	"source":   &models.Source{},
	"dgima":    &models.Dgima{},
	"aricha":   &models.Aricha{},
	"jobs":     &models.Job{},
	"files":    &models.File{},
	"cloud":    &models.Cloud{},
	"users":    &models.User{},
	"label":    &models.Label{},
}

var ids = map[string]string{
	"ingest":   "capture_id",
	"trimmer":  "trim_id",
	"products": "product_id",
	"state":    "state_id",
	"kmedia":   "kmedia_id",
	"archive":  "archive_id",
	"capture":  "capture_id",
	"convert":  "convert_id",
	"carbon":   "carbon_id",
	"insert":   "insert_id",
	"source":   "source_id",
	"dgima":    "dgima_id",
	"aricha":   "aricha_id",
	"jobs":     "job_id",
	"files":    "file_id",
	"cloud":    "cloud_id",
	"users":    "user_id",
	"label":    "id",
}

func V1GetRecordsByKV(c *gin.Context) {
	root := c.Params.ByName("root")
	key := c.Query("key")
	val := c.Query("value")
	t := list[root]
	if r, err := models.V1FindByKV(key, val, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func V2GetRecordsByKV(c *gin.Context) {
	root := c.Params.ByName("root")
	values := c.Request.URL.Query()
	t := list[root]
	if r, err := models.V2FindByKV(root, values, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetRecordsByJSON(c *gin.Context) {
	root := c.Params.ByName("root")
	prop := c.Params.ByName("prop")
	values := c.Request.URL.Query()
	t := list[root]
	if r, err := models.FindByJSON(root, prop, values, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetRecordByID(c *gin.Context) {
	root := c.Params.ByName("root")
	idVal := c.Params.ByName("id")
	idKey := ids[root]
	t := recd[root]
	if r, err := models.FindByID(idKey, idVal, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func PutRecord(c *gin.Context) {
	root := c.Params.ByName("root")
	t := recd[root]
	err := c.BindJSON(&t)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}
	err = models.CreateRecord(t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
		go SendMessage(root)
	}
}

func UpdateRecord(c *gin.Context) {
	root := c.Params.ByName("root")
	idKey := ids[root]
	idVal := c.Params.ByName("id")
	key := c.Params.ByName("key")

	// JSONB we take from body and simple value from option
	val := c.Query("value")
	var err error
	var t map[string]interface{}
	err = c.ShouldBindJSON(&t)

	if val == "" && err != nil {
		NewBadRequestError(err).Abort(c)
		return
	}

	// Ignore value option if body exist
	if err == nil {
		val = ""
		err = models.UpdateRecord(idKey, idVal, key, t, root)
	}

	if val != "" {
		err = models.UpdateRecord(idKey, idVal, key, val, root)
	}

	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
		go SendMessage(root)
	}
}

func UpdateJsonbRecord(c *gin.Context) {
	root := c.Params.ByName("root")
	idKey := ids[root]
	idVal := c.Params.ByName("id")
	key := c.Params.ByName("key")
	prop := c.Params.ByName("prop")

	// JSONB we take from body and simple value from option
	val := c.Query("value")
	var err error
	var t map[string]interface{}
	err = c.ShouldBindJSON(&t)

	if val == "" && err != nil {
		NewBadRequestError(err).Abort(c)
		return
	}

	// Ignore value option if body exist
	if err == nil {
		val = ""
		err = models.UpdateJSONB(idKey, idVal, key, t, root, prop)
	}

	if val != "" {
		err = models.UpdateJSONB(idKey, idVal, key, val, root, prop)
	}

	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
		go SendMessage(root)
	}
}

func RemoveRecord(c *gin.Context) {
	root := c.Params.ByName("root")
	idKey := ids[root]
	idVal := c.Params.ByName("id")
	t := recd[root]
	err := models.RemoveRecord(idKey, idVal, t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
		go SendMessage(root)
	}
}

// Trimmer

func GetTrimmed(c *gin.Context) {
	var t []models.Trimmer
	if r, err := models.FindTrimmed(t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

// Jobs

func GetJobs(c *gin.Context) {
	var t []models.Job
	if r, err := models.FindJobs(t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

// Dgima

func GetDgima(c *gin.Context) {
	var t []models.Dgima
	if r, err := models.FindDgima(t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

// Aricha

func GetAricha(c *gin.Context) {
	var t []models.Dgima
	if r, err := models.FindAricha(t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

// State

func GetStates(c *gin.Context) {
	if r, err := models.GetStates(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetStateByID(c *gin.Context) {
	id := c.Params.ByName("id")
	if r, err := models.GetState(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetStateByTag(c *gin.Context) {
	id := c.Params.ByName("id")
	tag := c.Params.ByName("tag")
	if r, err := models.GetStateByTag(id, tag); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetStateByProp(c *gin.Context) {
	id := c.Params.ByName("id")
	prop := c.Params.ByName("prop")
	if r, err := models.GetStateProp(id, prop); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func RemoveStateByProp(c *gin.Context) {
	id := c.Params.ByName("id")
	prop := c.Params.ByName("prop")
	err := models.RemoveStateProp(id, prop)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
		go SendMessage("langcheck")
	}
}

func PutStateByID(c *gin.Context) {
	var t models.State
	t.StateID = c.Params.ByName("id")
	t.Tag = c.Params.ByName("tag")
	err := c.BindJSON(&t.Data)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}

	err = models.CreateRecord(t)

	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func PutStateByProp(c *gin.Context) {
	var t map[string]interface{}
	err := c.ShouldBindJSON(&t)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}
	id := c.Params.ByName("id")
	prop := c.Params.ByName("prop")

	err = models.UpdateJSONB("state_id", id, "data", t, "state", prop)

	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}
