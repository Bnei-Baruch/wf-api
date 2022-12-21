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

func GetState(id string) (interface{}, error) {
	var state interface{}
	r := DB.Raw("SELECT data FROM state WHERE state_id = ?", id).Scan(&state)
	if r.Error != nil {
		return nil, r.Error
	}

	return state, nil
}

func GetStateProp(id string, prop string) (interface{}, error) {
	var state interface{}
	r := DB.Raw("SELECT data->>? FROM state WHERE state_id = ?", prop, id).Scan(&state)
	if r.Error != nil {
		return nil, r.Error
	}

	return state, nil
}

func RemoveStateProp(id string, prop string) error {
	r := DB.Exec("UPDATE state SET data = data - ? WHERE state_id = ?", prop, id)
	if r.Error != nil {
		return r.Error
	}

	return nil
}
