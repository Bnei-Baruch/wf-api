package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

func (Trimmer) TableName() string {
	return "trimmer"
}

type Trimmer struct {
	TrimID    string          `json:"trim_id" gorm:"primaryKey"`
	Date      string          `json:"date"`
	FileName  string          `json:"file_name"`
	Inpoints  pq.Float32Array `json:"inpoints" gorm:"type:numeric"`
	Outpoints pq.Float32Array `json:"outpoints" gorm:"type:numeric"`
	Parent    datatypes.JSON  `json:"parent" gorm:"type:jsonb"`
	Line      datatypes.JSON  `json:"line" gorm:"type:jsonb"`
	Original  datatypes.JSON  `json:"original" gorm:"type:jsonb"`
	Proxy     datatypes.JSON  `json:"proxy" gorm:"type:jsonb"`
	Wfstatus  datatypes.JSON  `json:"wfstatus" gorm:"type:jsonb"`
}

func FindTrimmed(t interface{}) (interface{}, error) {
	err := DB.Order("id").Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error
	return t, err
}
