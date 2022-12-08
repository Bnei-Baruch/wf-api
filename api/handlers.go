package api

import (
	"github.com/Bnei-Baruch/wf-api/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var u models.User
	if c.BindJSON(&u) == nil {
		models.DB.Create(&u)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
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

func GetIngest(c *gin.Context) {
	id := c.Params.ByName("id")
	var ingest models.Ingest
	if err := models.DB.Where("capture_id = ?", id).First(&ingest).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ingest)
	}
}

func FindTrimmer(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("value")
	var t []models.Trimmer
	if err := models.DB.Where(key+" LIKE ?", "%"+val+"%").Find(&t).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, t)
	}
}

func GetTrimmer(c *gin.Context) {
	id := c.Params.ByName("id")
	var t models.Trimmer
	if err := models.DB.Where("trim_id = ?", id).First(&t).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, t)
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

func PutTrimmer(c *gin.Context) {
	var t models.Trimmer
	if c.BindJSON(&t) == nil {
		models.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&t)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func TrimmerStatusValue(c *gin.Context) {
	id := c.Params.ByName("id")
	key := c.Params.ByName("key")
	val := c.Query("value")
	models.DB.Exec("UPDATE trimmer SET wfstatus = wfstatus || json_build_object($2::text, $3::bool)::jsonb WHERE trim_id=$1", id, key, val)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func RemoveTrimmer(c *gin.Context) {
	id := c.Params.ByName("id")
	var t models.Trimmer
	if err := models.DB.Unscoped().Where("trim_id = ?", id).Delete(&t).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
