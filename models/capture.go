package models

import "github.com/jackc/pgtype"

type Capture struct {
	ID        int          `json:"id" gorm:"autoIncrement"`
	CaptureID string       `json:"capture_id" gorm:"primaryKey"`
	Capsrc    string       `json:"capture_src"`
	Date      string       `json:"date"`
	StartName string       `json:"start_name"`
	StopName  string       `json:"stop_name"`
	Sha1      string       `json:"sha1"`
	Line      pgtype.JSONB `json:"line" gorm:"type:jsonb"`
	Original  pgtype.JSONB `json:"original" gorm:"type:jsonb"`
	Proxy     pgtype.JSONB `json:"proxy" gorm:"type:jsonb"`
	Wfstatus  pgtype.JSONB `json:"wfstatus" gorm:"type:jsonb"`
}
