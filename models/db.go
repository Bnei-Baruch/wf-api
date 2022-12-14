package models

import (
	"fmt"
	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Tabler interface {
	TableName() string
}

func InitDB() {
	log.Info("Setting up connection to DB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jerusalem",
		viper.GetString("db.host"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.name"),
		viper.GetString("db.port"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gorm_logrus.New().LogMode(logger.Info),
	})
	if err != nil {
		log.Infof("DB connection error: %s", err)
		return
	}
	DB.AutoMigrate()
}

func CreateRecord(s interface{}) error {
	r := DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(s)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func UpdateRecord(idKey string, idVal string, colKey string, colVal interface{}, table string) error {
	sqlCmd := "UPDATE " + table + " SET " + colKey + " = $2 WHERE " + idKey + "=$1"
	r := DB.Exec(sqlCmd, idVal, colVal)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func UpdateJSONB(idKey string, idVal string, propKey string, propVal interface{}, table string, prop string) error {
	var sqlCmd string
	var varType string
	if _, ok := propVal.(string); ok {
		if propVal == "true" || propVal == "false" {
			varType = "bool"
		} else {
			varType = "text"
		}
	} else {
		varType = "jsonb"
	}

	// TODO: Add other types

	sqlCmd = "UPDATE " + table + " SET " + prop + " = " + prop + " || json_build_object($1::text, $2::" + varType + ")::jsonb WHERE " + idKey + "=$3"

	r := DB.Exec(sqlCmd, propKey, propVal, idVal)

	if r.Error != nil {
		return r.Error
	}
	return nil
}

func RemoveRecord(idKey string, idVal string, s interface{}) error {
	r := DB.Unscoped().Where(idKey+" = ?", idVal).Delete(&s)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func FindByKV(key string, val string, s interface{}) (interface{}, error) {
	err := DB.Where(key+" LIKE ?", "%"+val+"%").Find(&s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func FindByID(key string, id string, s interface{}) (interface{}, error) {
	r := DB.Where(key+" = ?", id).First(&s)
	if r.Error != nil {
		return s, r.Error
	}
	return s, nil
}
