package models

type Carbon struct {
	ID        int     `json:"id" gorm:"autoIncrement"`
	CarbonID  string  `json:"carbon_id" gorm:"primaryKey"`
	SendID    string  `json:"send_id"`
	Date      string  `json:"date"`
	FileName  string  `json:"file_name"`
	Language  string  `json:"language"`
	Extension string  `json:"extension"`
	Duration  float32 `json:"duration"`
	Size      int64   `json:"size"`
	Sha1      string  `json:"sha1"`
}
