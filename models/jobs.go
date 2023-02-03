package models

import (
	"github.com/jackc/pgtype"
)

type Job struct {
	ID       int          `json:"id" gorm:"autoIncrement"`
	JobID    string       `json:"job_id" gorm:"primaryKey"`
	Date     string       `json:"date"`
	FileName string       `json:"file_name"`
	JobName  string       `json:"job_name"`
	JobType  string       `json:"job_type"`
	Parent   pgtype.JSONB `json:"parent" gorm:"type:jsonb"`
	Line     pgtype.JSONB `json:"line" gorm:"type:jsonb"`
	Original pgtype.JSONB `json:"original" gorm:"type:jsonb"`
	Proxy    pgtype.JSONB `json:"proxy" gorm:"type:jsonb"`
	Product  pgtype.JSONB `json:"product" gorm:"type:jsonb"`
	Wfstatus pgtype.JSONB `json:"wfstatus" gorm:"type:jsonb"`
}

func FindJobs(t interface{}) (interface{}, error) {
	err := DB.Where("wfstatus ->> 'removed' = ?", "false").Find(&t).Error
	return t, err
}
