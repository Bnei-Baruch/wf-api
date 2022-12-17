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
}

var recd = map[string]interface{}{
	"ingest":   &models.Ingest{},
	"trimmer":  &models.Trimmer{},
	"products": &models.Product{},
	"state":    &models.State{},
	"kmedia":   &models.Kmedia{},
}

var ids = map[string]string{
	"ingest":   "capture_id",
	"trimmer":  "trim_id",
	"products": "product_id",
	"state":    "state_id",
	"kmedia":   "kmedia_id",
}

func GetRecordsByKV(c *gin.Context) {
	root := c.Params.ByName("root")
	key := c.Query("key")
	val := c.Query("value")
	t := list[root]
	if r, err := models.FindByKV(key, val, t); err != nil {
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
	err = models.CreateRecord(&t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
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
	err = c.BindJSON(&t)

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
	err = c.BindJSON(&t)

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
	}
}

func RemoveRecord(c *gin.Context) {
	root := c.Params.ByName("root")
	idKey := ids[root]
	idVal := c.Params.ByName("id")
	t := recd[root]
	err := models.RemoveRecord(idKey, idVal, &t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
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
