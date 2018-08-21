package model

type Tag struct {
	BaseModel
	Name					string		`json:"name" gorm:"type:varchar(30);not null;unique_index" validate:"required,min=2,max=10"`
	Description		string		`json:"desc" gorm:"type:varchar(150)" validate:"max=50"`
	AuthorID			uint			`json:"authorID" validate:"required"`
	Total					uint			`json:"total" gorm:"-"`		// 标签使用数量
}
