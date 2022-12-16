package models

import "github.com/jackc/pgtype"

type User struct {
	ID         int          `json:"id" gorm:"autoIncrement"`
	UserID     string       `json:"user_id" gorm:"primaryKey"`
	FirstName  string       `json:"firstName"`
	LastName   string       `json:"lastName"`
	Email      string       `json:"email"`
	Properties pgtype.JSONB `json:"properties" gorm:"type:jsonb"`
	Role       string       `json:"role"`
}
