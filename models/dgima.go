package models

import (
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
)

func (Dgima) TableName() string {
	return "dgima"
}

type Dgima struct {
	ID        int             `json:"id" gorm:"autoIncrement"`
	TrimID    string          `json:"dgima_id" gorm:"primaryKey"`
	Date      string          `json:"date"`
	FileName  string          `json:"file_name"`
	Inpoints  pq.Float32Array `json:"inpoints" gorm:"type:numeric"`
	Outpoints pq.Float32Array `json:"outpoints" gorm:"type:numeric"`
	Parent    pgtype.JSONB    `json:"parent"`
	Line      pgtype.JSONB    `json:"line" gorm:"type:jsonb"`
	Original  pgtype.JSONB    `json:"original" gorm:"type:jsonb"`
	Proxy     pgtype.JSONB    `json:"proxy" gorm:"type:jsonb"`
	Wfstatus  pgtype.JSONB    `json:"wfstatus" gorm:"type:jsonb"`
}
