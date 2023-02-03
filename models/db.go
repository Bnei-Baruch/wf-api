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
	"net/url"
	"regexp"
	"strconv"
	"strings"
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

func RemoveRecord(idKey string, idVal string, s interface{}) error {
	r := DB.Unscoped().Where(idKey+" = ?", idVal).Delete(&s)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func FindByID(key string, id string, s interface{}) (interface{}, error) {
	r := DB.Where(key+" = ?", id).First(&s)
	if r.Error != nil {
		return s, r.Error
	}
	return s, nil
}

func V1FindByKV(key string, val string, s interface{}) (interface{}, error) {
	err := DB.Debug().Where(key+" LIKE ?", "%"+val+"%").Find(&s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func V2FindByKV(table string, values url.Values, t interface{}) (interface{}, error) {
	var where []string
	sqlStatement := "SELECT * FROM " + table
	limit := "100"
	offset := "0"
	i := 0
	chk, _ := regexp.MatchString(`^(trimmer|aricha|dgima|ingest|capture)$`, table)

	for k, v := range values {
		if k == "limit" {
			limit = v[0]
			continue
		}
		if k == "offset" {
			offset = v[0]
			continue
		}

		// FIXME: It's files endpoint compb
		if k == "archive" || k == "mdb" {
			if i == 0 {
				sqlStatement = sqlStatement + ` WHERE ` + fmt.Sprintf(`properties['%s'] = '%s'`, k, v[0])
				i += 1
				continue
			} else {
				where = append(where, fmt.Sprintf(`properties['%s'] = '%s'`, k, v[0]))
				i += 1
				continue
			}
		}

		// FIXME: we need to move json to first tree level, write now sha option must first in chk root
		if chk && k == "sha1" && i == 0 {
			sqlStatement = sqlStatement + ` WHERE ` + fmt.Sprintf(`original['format']['sha1'] = '"%s"'`, v[0])
			i += 1
			continue
		}

		if i == 0 {
			sqlStatement = sqlStatement + ` WHERE ` + fmt.Sprintf(`"%s" = '%s'`, k, v[0])
		} else {
			where = append(where, fmt.Sprintf(`"%s" = '%s'`, k, v[0]))
		}

		i += 1
	}

	if len(where) > 0 {
		sqlStatement = sqlStatement + ` AND ` + strings.Join(where, " AND ")
	}

	sqlStatement = sqlStatement + fmt.Sprintf(` ORDER BY id DESC LIMIT %s OFFSET %s`, limit, offset)

	r := DB.Raw(sqlStatement).Scan(&t)

	if r.Error != nil {
		return nil, r.Error
	}

	return t, nil
}

func FindByJSON(table string, prop string, values url.Values, t interface{}) (interface{}, error) {
	var where []string
	var q string
	sqlStatement := "SELECT * FROM " + table
	limit := "100"
	offset := "0"
	i := 0

	for k, v := range values {
		if k == "limit" {
			limit = v[0]
			continue
		}
		if k == "offset" {
			offset = v[0]
			continue
		}
		if _, err := strconv.ParseBool(v[0]); err == nil {
			q = fmt.Sprintf(`%s['%s'] = '%s'`, prop, k, v[0])
		} else {
			q = fmt.Sprintf(`%s['%s'] = '"%s"'`, prop, k, v[0])
		}

		// TODO: check parse int

		if i == 0 {
			sqlStatement = sqlStatement + ` WHERE ` + q
		} else {
			where = append(where, q)
		}

		i += 1
	}

	if len(where) > 0 {
		sqlStatement = sqlStatement + ` AND ` + strings.Join(where, " AND ")
	}

	sqlStatement = sqlStatement + fmt.Sprintf(` ORDER BY id DESC LIMIT %s OFFSET %s`, limit, offset)

	r := DB.Raw(sqlStatement).Scan(&t)

	if r.Error != nil {
		return nil, r.Error
	}

	return t, nil
}

func UpdateJSONB(idKey string, idVal string, prop string, propVal interface{}, table string, propKey string) error {
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
