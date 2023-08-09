package models

func (Archive) TableName() string {
	return "archive"
}

type Archive struct {
	ArchiveID string `json:"archive_id" gorm:"primaryKey"`
	Date      string `json:"date"`
	FileName  string `json:"file_name"`
	Language  string `json:"language"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
	Sha1      string `json:"sha1"`
	SendID    string `json:"send_id"`
	Source    string `json:"source"`
}
