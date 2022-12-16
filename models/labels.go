package models

type Label struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Date         string `json:"date"`
	Lecturer     string `json:"lecturer"`
	Subject      string `json:"subject"`
	Language     string `json:"language"`
	Location     string `json:"location"`
	ContentType  string `json:"content_type"`
	CasseteType  string `json:"cassete_type"`
	Mof          string `json:"mof"`
	Duration     string `json:"duration"`
	ArchivePlace string `json:"archive_place"`
	Comments     string `json:"comments"`
	BarCode      string `json:"bar_code"`
}
