package models

import "gorm.io/datatypes"

func (Convert) TableName() string {
	return "convert"
}

type Convert struct {
	ConvertID string         `json:"convert_id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Date      string         `json:"date"`
	Progress  string         `json:"progress"`
	State     string         `json:"state"`
	Timestamp string         `json:"timestamp"`
	Langcheck datatypes.JSON `json:"langcheck" gorm:"type:jsonb"`
}
