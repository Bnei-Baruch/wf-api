package models

func (Kmedia) TableName() string {
	return "kmedia"
}

type Kmedia struct {
	ID        int    `json:"id" gorm:"autoIncrement"`
	KmediaID  string `json:"kmedia_id" gorm:"primaryKey"`
	Date      string `json:"date"`
	FileName  string `json:"file_name"`
	Language  string `json:"language"`
	Extension string `json:"extension"`
	Size      int64  `json:"size"`
	Sha1      string `json:"sha1"`
	Pattern   string `json:"pattern"`
	SendID    string `json:"send_id"`
	Source    string `json:"source"`
}
