package models

import (
	"fmt"
	"gorm.io/datatypes"
	"net/url"
	"strings"
)

type Job struct {
	JobID    string         `json:"job_id" gorm:"primaryKey"`
	Date     string         `json:"date"`
	FileName string         `json:"file_name"`
	JobName  string         `json:"job_name"`
	JobType  string         `json:"job_type"`
	Parent   datatypes.JSON `json:"parent" gorm:"type:jsonb"`
	Line     datatypes.JSON `json:"line" gorm:"type:jsonb"`
	Original datatypes.JSON `json:"original" gorm:"type:jsonb"`
	Proxy    datatypes.JSON `json:"proxy" gorm:"type:jsonb"`
	Product  datatypes.JSON `json:"product" gorm:"type:jsonb"`
	Wfstatus datatypes.JSON `json:"wfstatus" gorm:"type:jsonb"`
}

func FindJobs(t interface{}) (interface{}, error) {
	err := DB.Order("id").Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error
	return t, err
}

func FindJobsByUserID(values url.Values, t interface{}) (interface{}, error) {
	var where []string
	sqlStatement := `SELECT * FROM jobs WHERE wfstatus['removed'] = 'false'`
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
		if k == "doers" {
			where = append(where, fmt.Sprintf(`parent['%s'] ? '"%s"'`, k, v[0]))
			continue
		}
		where = append(where, fmt.Sprintf(`"%s" = '%s'`, k, v[0]))
	}

	if len(where) > 0 {
		sqlStatement = sqlStatement + ` AND ` + strings.Join(where, " AND ")
	}

	sqlStatement = sqlStatement + fmt.Sprintf(` ORDER BY job_id DESC LIMIT %s OFFSET %s`, limit, offset)

	r := DB.Raw(sqlStatement).Scan(&t)

	if r.Error != nil {
		return nil, r.Error
	}

	return t, nil
}
