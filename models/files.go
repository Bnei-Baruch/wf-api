package models

import "gorm.io/datatypes"

type File struct {
	FileID     string         `json:"file_id" gorm:"primaryKey"`
	Date       string         `json:"date"`
	FileName   string         `json:"file_name"`
	Extension  string         `json:"extension"`
	Size       int64          `json:"size"`
	Sha1       string         `json:"sha1"`
	FileType   string         `json:"file_type"`
	Language   string         `json:"language"`
	MimeType   string         `json:"mime_type"`
	UID        string         `json:"uid"`
	WID        string         `json:"wid"`
	Properties datatypes.JSON `json:"properties" gorm:"type:jsonb"`
	ProductID  string         `json:"product_id"`
	MediaInfo  datatypes.JSON `json:"media_info" gorm:"type:jsonb"`
}
