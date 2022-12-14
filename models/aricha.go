package models

import (
	"github.com/jackc/pgtype"
)

func (Aricha) TableName() string {
	return "aricha"
}

type Aricha struct {
	ID       int          `json:"id" gorm:"autoIncrement"`
	TrimID   string       `json:"aricha_id" gorm:"primaryKey"`
	Date     string       `json:"date"`
	FileName string       `json:"file_name"`
	Parent   pgtype.JSONB `json:"parent"`
	Line     pgtype.JSONB `json:"line" gorm:"type:jsonb"`
	Original pgtype.JSONB `json:"original" gorm:"type:jsonb"`
	Proxy    pgtype.JSONB `json:"proxy" gorm:"type:jsonb"`
	Wfstatus pgtype.JSONB `json:"wfstatus" gorm:"type:jsonb"`
}
