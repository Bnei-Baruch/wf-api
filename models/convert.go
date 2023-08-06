package models

import "gorm.io/datatypes"

func (Convert) TableName() string {
	return "convert"
}

type Convert struct {
	ID        int            `json:"id" gorm:"autoIncrement"`
	ConvertID string         `json:"convert_id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Date      string         `json:"date"`
	Progress  string         `json:"progress"`
	State     string         `json:"state"`
	Timestamp string         `json:"timestamp"`
	Langcheck datatypes.JSON `json:"langcheck" gorm:"type:jsonb"`
}
