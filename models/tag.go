package models

import (
	sql "groot/db"
)

type Tag struct {
	BaseModel
	Name					string		`json:"name" gorm:"type:varchar(30);not null;unique_index" validate:"required,min=2,max=10"`
	Description		string		`json:"desc" gorm:"type:varchar(150)" validate:"max=50"`
	AuthorID			uint			`json:"authorID" validate:"required"`
	Total					int				`json:"total" gorm:"-"`		// 标签使用数量
}

func (tag *Tag) Save() error {
	return sql.DB.Create(tag).Error
}

func (tag *Tag) isExist() bool {
	return !sql.DB.Where("name = ?", tag.Name).Find(tag).RecordNotFound()
}

func (tag *Tag) FindAndSave() error {
	isExist := tag.isExist()
	var err error
	if !isExist {
		err = tag.Save()
	}

	return err
}
