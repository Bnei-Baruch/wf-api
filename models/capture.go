package models

import "gorm.io/datatypes"

func (Capture) TableName() string {
	return "capture"
}

type Capture struct {
	CaptureID  string         `json:"capture_id" gorm:"primaryKey"`
	CaptureSrc string         `json:"capture_src"`
	Date       string         `json:"date"`
	StartName  string         `json:"start_name"`
	StopName   string         `json:"stop_name"`
	Sha1       string         `json:"sha1"`
	Line       datatypes.JSON `json:"line" gorm:"type:jsonb"`
	Original   datatypes.JSON `json:"original" gorm:"type:jsonb"`
	Proxy      datatypes.JSON `json:"proxy" gorm:"type:jsonb"`
	Wfstatus   datatypes.JSON `json:"wfstatus" gorm:"type:jsonb"`
}
