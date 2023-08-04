package models

import (
	"errors"
	"github.com/jackc/pgtype"
)
import "encoding/json"

func (Source) TableName() string {
	return "source"
}

type Source struct {
	ID       int          `json:"id" gorm:"autoIncrement"`
	SourceID string       `json:"source_id" gorm:"primaryKey"`
	Date     string       `json:"date"`
	FileName string       `json:"file_name"`
	Sha1     string       `json:"sha1"`
	Line     pgtype.JSONB `json:"line" gorm:"type:jsonb"`
	Source   pgtype.JSONB `json:"source" gorm:"type:jsonb"`
	Wfstatus pgtype.JSONB `json:"wfstatus" gorm:"type:jsonb"`
}

func GetSourceByUID(uid string) (interface{}, error) {
	t := map[string]interface{}{}

	r := DB.Raw("SELECT source['kmedia'] FROM source WHERE source->'kmedia'->>'file_uid' = ?", uid).Scan(&t)
	if r.Error != nil {
		return nil, r.Error
	}

	if t["source"] == nil {
		err := errors.New("not found")
		return nil, err
	}

	j := t["source"].(string)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(j), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
