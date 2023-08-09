package models

import "gorm.io/datatypes"

func (Cloud) TableName() string {
	return "cloud"
}

type Cloud struct {
	OID        string         `json:"oid" gorm:"primaryKey"`
	Date       string         `json:"date"`
	Name       string         `json:"name"`
	Type       string         `json:"type"`
	Extension  string         `json:"extension"`
	Language   string         `json:"language"`
	Source     string         `json:"source"`
	WID        string         `json:"wid"`
	UID        string         `json:"uid"`
	Pattern    string         `json:"pattern"`
	Properties datatypes.JSON `json:"properties" gorm:"type:jsonb"`
	Url        string         `json:"url"`
}
