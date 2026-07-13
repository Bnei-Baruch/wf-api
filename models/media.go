package models

import (
	"gorm.io/datatypes"
)

func (Media) TableName() string {
	return "media"
}

type Media struct {
	MediaID  string         `json:"media_id" gorm:"primaryKey"`
	Date     string         `json:"date"`
	FileName string         `json:"file_name"`
	FilePath string         `json:"file_path"`
	Size     int64          `json:"size"`
	Sha1     string         `json:"sha1"`
	UID      string         `json:"uid"`
	LabelID  string         `json:"label_id"`
	Line     datatypes.JSON `json:"line" gorm:"type:jsonb"`
	Extra    datatypes.JSON `json:"extra" gorm:"type:jsonb"`
	Wfstatus datatypes.JSON `json:"wfstatus" gorm:"type:jsonb"`
}

func FindMedia(t interface{}) (interface{}, error) {
	err := DB.Order("id").Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error
	return t, err
}
