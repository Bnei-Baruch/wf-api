package models

import "gorm.io/datatypes"

func (Insert) TableName() string {
	return "insert"
}

type Insert struct {
	InsertID   string         `json:"insert_id" gorm:"primaryKey"`
	InsertName string         `json:"insert_name"`
	Date       string         `json:"date"`
	FileName   string         `json:"file_name"`
	Extension  string         `json:"extension"`
	Size       int64          `json:"size"`
	Sha1       string         `json:"sha1"`
	Language   string         `json:"language"`
	InsertType string         `json:"insert_type"`
	SendID     string         `json:"send_id"`
	UploadType string         `json:"upload_type"`
	Line       datatypes.JSON `json:"line" gorm:"type:jsonb"`
}
