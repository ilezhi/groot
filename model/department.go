package model

type Department struct {
	BaseModel
	Code					int					`json:"code" gorm:"type:tinyint"`
	Name					string			`json:"name" gorm:"size:30"`
	AliasName 		string			`json:"aliasName" gorm:"size:60"`
	Description		string			`json:"description"`
}
