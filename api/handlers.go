package api

import (
	"github.com/Bnei-Baruch/wf-api/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var u models.User
	if c.BindJSON(&u) == nil {
		models.DB.Create(&u)
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func GetUsers(c *gin.Context) {
	var u []models.User
	if err := models.DB.Find(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func GetUser(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}
	var u models.User
	if err := models.DB.Where("id = ?", id).First(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, u)
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

func UpdateIngestState(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	val := c.Query("value")
	err := models.UpdateState("capture_id", id, key, val, "ingest")
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
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func UpdateTrimmerState(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	val := c.Query("value")
	err := models.UpdateState("trim_id", id, key, val, "trimmer")
	if err != nil {
		NewInternalError(err).Abort(c)
	} else {
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
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func GetTrimmed(c *gin.Context) {
	var t []models.Trimmer
	if err := models.DB.Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, t)
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
