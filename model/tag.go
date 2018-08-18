package model

type Tag struct {
	BaseModel
	Name					string		`gorm:"type:varchar(30);not null;unique_index"`
	Description		string		`gorm:"type:varchar(150)"`
	AuthorID			uint
	Total					uint			`gorm:"-"`		// 标签使用数量
}
