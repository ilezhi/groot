package models

import (
	sql "groot/db"
)

// 类别: 将收藏的topic 进行分类
type Category struct {
	BaseModel
	Name				string			`json:"name" gorm:"size:30;not null"`
	UserID			uint				`json:"userID"`
	Count				int					`json:"count" gorm:"-"`
}

func (c *Category) IsExist() bool {
	return !sql.DB.Where(c).Find(c).RecordNotFound()
}

func (c *Category) Save() error {
	return sql.DB.Create(c).Error
}
