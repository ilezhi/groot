package models

import "time"

type BaseModel struct {
	ID					uint			`json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt		time.Time	`json:"createdAt"`
	UpdatedAt		time.Time `json:"updatedAt"`
}
