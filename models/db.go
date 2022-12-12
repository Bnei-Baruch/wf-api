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
	r := DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&s)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func UpdateState(idKey string, idVal string, key string, val string, table string) error {
	r := DB.Exec("UPDATE $5 SET wfstatus = wfstatus || json_build_object($3::text, $4::bool)::jsonb WHERE $1=$2", idKey, idVal, key, val, table)
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
	err := DB.Where(key+" = ?", id).Find(&s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}
