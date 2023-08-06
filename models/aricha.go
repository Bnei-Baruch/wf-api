package models

import "gorm.io/datatypes"

func (Aricha) TableName() string {
	return "aricha"
}

type Aricha struct {
	ID       int            `json:"id" gorm:"autoIncrement"`
	ArichaID string         `json:"aricha_id" gorm:"primaryKey"`
	Date     string         `json:"date"`
	FileName string         `json:"file_name"`
	Parent   datatypes.JSON `json:"parent" gorm:"type:jsonb"`
	Line     datatypes.JSON `json:"line" gorm:"type:jsonb"`
	Original datatypes.JSON `json:"original" gorm:"type:jsonb"`
	Proxy    datatypes.JSON `json:"proxy" gorm:"type:jsonb"`
	Wfstatus datatypes.JSON `json:"wfstatus" gorm:"type:jsonb"`
}

func FindAricha(t interface{}) (interface{}, error) {
	err := DB.Order("id").Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error
	return t, err
}
