package models

import "github.com/jackc/pgtype"

func (Source) TableName() string {
	return "source"
}

type Source struct {
	ID       int          `json:"id" gorm:"autoIncrement"`
	SourceID string       `json:"source_id" gorm:"primaryKey"`
	Date     string       `json:"date"`
	FileName string       `json:"file_name"`
	Sha1     string       `json:"sha1"`
	Line     pgtype.JSONB `json:"line" gorm:"type:jsonb"`
	Source   pgtype.JSONB `json:"source" gorm:"type:jsonb"`
	Wfstatus pgtype.JSONB `json:"wfstatus" gorm:"type:jsonb"`
}

func GetSourceByUID(t interface{}, uid string) (interface{}, error) {
	r := DB.Raw("SELECT source->'kmedia' FROM source WHERE source->'kmedia'->>'file_uid' =  ?", uid).Scan(&t)
	if r.Error != nil {
		return nil, r.Error
	}

	return t, nil
}
