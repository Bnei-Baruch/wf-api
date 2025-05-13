package models

func (Backup) TableName() string {
	return "backup"
}

type Backup struct {
	ID       uint   `json:"id" gorm:"autoIncrement"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	FileSha1 string `json:"file_sha1" gorm:"primaryKey"`
	FileSize int64  `json:"file_size"`
	FileDate int64  `json:"file_date"`
}
