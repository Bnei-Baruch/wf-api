package models

type State struct {
	ID      int         `json:"id" gorm:"autoIncrement"`
	StateID string      `json:"state_id" gorm:"primaryKey"`
	Data    interface{} `json:"data"`
	Tag     string      `json:"tag"`
}
