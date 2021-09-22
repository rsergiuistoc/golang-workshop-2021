package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Todo struct {
	ID 				uuid.UUID 	`gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	Title			string		`gorm:"size:150;not null" json:"title"`
	Description 	string 		`gorm:"type:text;not null" json:"description"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	UserID 			uuid.UUID	`json:"-"`
}