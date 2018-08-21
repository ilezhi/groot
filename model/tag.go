package model

type Tag struct {
	BaseModel
	Name					string		`json:"name" gorm:"type:varchar(30);not null;unique_index"`
	Description		string		`json:"desc" gorm:"type:varchar(150)"`
	AuthorID			uint			`json:"authorID"`
	Total					uint			`json:"total" gorm:"-"`		// 标签使用数量
}
