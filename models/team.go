package models

type Team struct {
	BaseModel
	Name					string			`json:"name" gorm:"size:30"`
	Description		string			`json:"description"`
}
