package models

import "github.com/jackc/pgtype"

func (State) TableName() string {
	return "state"
}

type State struct {
	ID      int          `json:"id" gorm:"autoIncrement"`
	StateID string       `json:"state_id" gorm:"primaryKey"`
	Data    pgtype.JSONB `json:"data" gorm:"type:jsonb"`
	Tag     string       `json:"tag"`
}
