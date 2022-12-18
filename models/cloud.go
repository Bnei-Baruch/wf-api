package models

import "github.com/jackc/pgtype"

func (Cloud) TableName() string {
	return "cloud"
}

type Cloud struct {
	ID         int          `json:"id" gorm:"autoIncrement"`
	OID        string       `json:"oid" gorm:"primaryKey"`
	Date       string       `json:"date"`
	Name       string       `json:"name"`
	Type       string       `json:"type"`
	Extension  string       `json:"extension"`
	Language   string       `json:"language"`
	Source     string       `json:"source"`
	WID        string       `json:"wid"`
	UID        string       `json:"uid"`
	Pattern    string       `json:"pattern"`
	Properties pgtype.JSONB `json:"properties" gorm:"type:jsonb"`
	Url        string       `json:"url"`
}
