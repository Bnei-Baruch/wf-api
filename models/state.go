package models

import (
	"encoding/json"
	"github.com/jackc/pgtype"
	"gorm.io/gorm/clause"
)

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
	var s State
	r := DB.Raw("SELECT data FROM state WHERE state_id = ?", id).Scan(&s)
	if r.Error != nil {
		return nil, r.Error
	}

	return s.Data, nil
}

func GetStateProp(id string, prop string) (interface{}, error) {
	var s []byte
	var data interface{}
	r := DB.Raw("SELECT data->>? FROM state WHERE state_id = ?", prop, id).Row()
	r.Scan(&s)
	_ = json.Unmarshal(s, &data)

	return data, nil
}

func RemoveStateProp(id string, prop string) error {
	r := DB.Exec("UPDATE state SET data = data - ? WHERE state_id = ?", prop, id)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func GetStateByTag(id string, tag string) (map[string]interface{}, error) {
	rows, err := DB.Model(&State{}).Where("tag = ?", tag).Select("id, state_id, data").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	states := make(map[string]interface{})

	for rows.Next() {
		var s State
		var o map[string]interface{}
		var obj []byte
		if err := rows.Scan(&s.ID, &s.StateID, &obj); err != nil {
			return nil, err
		}
		json.Unmarshal(obj, &o)
		states[s.StateID] = o
	}

	return states, nil
}

func GetStates() ([]State, error) {
	rows, err := DB.Model(&State{}).Select("id, state_id, data, tag").Order("tag").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []State{}

	for rows.Next() {
		var s State
		var obj []byte
		if err := rows.Scan(&s.ID, &s.StateID, &obj, &s.Tag); err != nil {
			return nil, err
		}
		json.Unmarshal(obj, &s.Data)
		states = append(states, s)
	}

	return states, nil
}

func PutStateByID(s interface{}) error {
	r := DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(s)
	if r.Error != nil {
		return r.Error
	}
	return nil
}
