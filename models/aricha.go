package models

import (
	"github.com/jackc/pgtype"
)

func (Aricha) TableName() string {
	return "aricha"
}

type Aricha struct {
	ID       int          `json:"id" gorm:"autoIncrement"`
	ArichaID string       `json:"aricha_id" gorm:"primaryKey"`
	Date     string       `json:"date"`
	FileName string       `json:"file_name"`
	Parent   pgtype.JSONB `json:"parent" gorm:"type:jsonb"`
	Line     pgtype.JSONB `json:"line" gorm:"type:jsonb"`
	Original pgtype.JSONB `json:"original" gorm:"type:jsonb"`
	Proxy    pgtype.JSONB `json:"proxy" gorm:"type:jsonb"`
	Wfstatus pgtype.JSONB `json:"wfstatus" gorm:"type:jsonb"`
}

func FindAricha(t interface{}) (interface{}, error) {
	err := DB.Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error
	return t, err
}
