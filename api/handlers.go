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

func PostRecordJSON(c *gin.Context) {
	root := c.Params.ByName("root")
	idKey := ids[root]
	idVal := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var t map[string]interface{}
	err := c.BindJSON(&t)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}
	err = models.UpdateRecord(idKey, idVal, key, t, root)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func UpdateRecordState(c *gin.Context) {
	root := c.Params.ByName("root")
	idKey := ids[root]
	idVal := c.Params.ByName("id")
	key := c.Params.ByName("key")
	st := c.Params.ByName("st")
	val := c.Query("value")
	err := models.UpdateJSONB(idKey, idVal, key, val, root, st)
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

// Ingest

func GetIngestByKV(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("value")
	var t []models.Ingest
	if r, err := models.FindByKV(key, val, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetIngestByID(c *gin.Context) {
	id := c.Params.ByName("id")
	key := "capture_id"
	var t models.Ingest
	if r, err := models.FindByID(key, id, &t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func PutIngest(c *gin.Context) {
	var t models.Ingest
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

func PostIngestJSON(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var t map[string]interface{}
	err := c.BindJSON(&t)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}
	err = models.UpdateRecord("capture_id", id, key, t, "ingest")
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func UpdateIngestState(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	st := c.Params.ByName("st")
	val := c.Query("value")
	err := models.UpdateJSONB("capture_id", id, key, val, "ingest", st)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func RemoveIngest(c *gin.Context) {
	id := c.Params.ByName("id")
	var t models.Ingest
	err := models.RemoveRecord("capture_id", id, &t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

// Trimmer

func GetTrimmerByKV(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("value")
	var t []models.Trimmer
	if r, err := models.FindByKV(key, val, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetTrimmerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	key := "trim_id"
	var t models.Trimmer
	if r, err := models.FindByID(key, id, &t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func PutTrimmer(c *gin.Context) {
	var t models.Trimmer
	err := c.BindJSON(&t)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}
	err = models.CreateRecord(&t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		go SendMessage("trim")
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func PostTrimmerJSON(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var t map[string]interface{}
	err := c.BindJSON(&t)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}
	err = models.UpdateRecord("trim_id", id, key, t, "trimmer")
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		go SendMessage("trim")
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func UpdateTrimmerState(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	st := c.Params.ByName("st")
	val := c.Query("value")
	err := models.UpdateJSONB("trim_id", id, key, val, "trimmer", st)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		go SendMessage("trim")
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func RemoveTrimmer(c *gin.Context) {
	id := c.Params.ByName("id")
	var t models.Trimmer
	err := models.RemoveRecord("trim_id", id, &t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		go SendMessage("trim")
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func GetTrimmed(c *gin.Context) {
	var t []models.Trimmer
	if r, err := models.FindTrimmed(t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

// Kmedia

func GetKmediaByKV(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("value")
	var t []models.Kmedia
	if r, err := models.FindByKV(key, val, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetKmediaByID(c *gin.Context) {
	id := c.Params.ByName("id")
	key := "kmedia_id"
	var t models.Kmedia
	if r, err := models.FindByID(key, id, &t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func PutKmedia(c *gin.Context) {
	var t models.Kmedia
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

func RemoveKmedia(c *gin.Context) {
	id := c.Params.ByName("id")
	var t models.Kmedia
	err := models.RemoveRecord("kmedia_id", id, &t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

// Products

func GetProductsByDF(c *gin.Context) {
	values := c.Request.URL.Query()
	var t []models.Product
	if r, err := models.FindByDF(values, t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	key := "product_id"
	var t models.Product
	if r, err := models.FindByID(key, id, &t); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func PutProduct(c *gin.Context) {
	var t models.Product
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

func UpdateProductState(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	st := c.Params.ByName("st")
	val := c.Query("value")
	err := models.UpdateJSONB("product_id", id, key, val, "products", st)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func PostProductJSON(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	var t map[string]interface{}
	err := c.BindJSON(&t)
	if err != nil {
		NewBadRequestError(err).Abort(c)
	}
	err = models.UpdateRecord("product_id", id, key, t, "products")
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func RemoveProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var t models.Product
	err := models.RemoveRecord("product_id", id, &t)
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
		go SendMessage("trim")
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}
