package models

import "gorm.io/datatypes"

type User struct {
	UserID     string         `json:"user_id" gorm:"primaryKey"`
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	Email      string         `json:"email"`
	Properties datatypes.JSON `json:"properties" gorm:"type:jsonb"`
	Role       string         `json:"role"`
}
