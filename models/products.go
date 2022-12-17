package models

import (
	"fmt"
	"github.com/jackc/pgtype"
	"net/url"
	"strings"
)

type Product struct {
	ID          int          `json:"id" gorm:"autoIncrement"`
	ProductID   string       `json:"product_id" gorm:"primaryKey"`
	Date        string       `json:"date"`
	Language    string       `json:"language"`
	Pattern     string       `json:"pattern"`
	TypeID      string       `json:"type_id"`
	ProductName string       `json:"product_name"`
	ProductType string       `json:"product_type"`
	I18n        pgtype.JSONB `json:"i18n" gorm:"type:jsonb"`
	Parent      pgtype.JSONB `json:"parent" gorm:"type:jsonb"`
	Line        pgtype.JSONB `json:"line" gorm:"type:jsonb"`
	Properties  pgtype.JSONB `json:"properties" gorm:"type:jsonb"`
	FilmDate    string       `json:"film_date"`
}

func FindByDF(values url.Values, t interface{}) (interface{}, error) {
	var where []string
	sqlStatement := `SELECT * FROM products WHERE properties['removed'] = 'false'`
	limit := "10"
	offset := "0"

	for k, v := range values {
		if k == "limit" {
			limit = v[0]
			continue
		}
		if k == "offset" {
			offset = v[0]
			continue
		}
		if k == "collection_uid" {
			where = append(where, fmt.Sprintf(`line['%s'] = '"%s"'`, k, v[0]))
			continue
		}
		where = append(where, fmt.Sprintf(`"%s" = '%s'`, k, v[0]))
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
