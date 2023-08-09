package models

import "gorm.io/datatypes"

type Job struct {
	JobID    string         `json:"job_id" gorm:"primaryKey"`
	Date     string         `json:"date"`
	FileName string         `json:"file_name"`
	JobName  string         `json:"job_name"`
	JobType  string         `json:"job_type"`
	Parent   datatypes.JSON `json:"parent" gorm:"type:jsonb"`
	Line     datatypes.JSON `json:"line" gorm:"type:jsonb"`
	Original datatypes.JSON `json:"original" gorm:"type:jsonb"`
	Proxy    datatypes.JSON `json:"proxy" gorm:"type:jsonb"`
	Product  datatypes.JSON `json:"product" gorm:"type:jsonb"`
	Wfstatus datatypes.JSON `json:"wfstatus" gorm:"type:jsonb"`
}

func FindJobs(t interface{}) (interface{}, error) {
	err := DB.Order("id").Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error
	return t, err
}
