package models

import "github.com/jackc/pgtype"

func (Convert) TableName() string {
	return "convert"
}

type Convert struct {
	ID        int          `json:"id" gorm:"autoIncrement"`
	ConvertID string       `json:"convert_id" gorm:"primaryKey"`
	Name      string       `json:"name"`
	Date      string       `json:"date"`
	Progress  string       `json:"progress"`
	State     string       `json:"state"`
	Timestamp string       `json:"timestamp"`
	Langcheck pgtype.JSONB `json:"langcheck" gorm:"type:jsonb"`
}
