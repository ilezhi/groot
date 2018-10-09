package models

import (
	// "github.com/jinzhu/gorm"
	// "gopkg.in/go-playground/validator.v9"
	sql "groot/db"
)

// 收藏
type Favor struct {
	BaseModel
	UserID			uint					`json:"userID" gorm:"index"`
	TopicID			uint					`json:"topicID" gorm:"index"`
	CategoryID	uint					`json:"categoryID"`
}

func (favor *Favor) IsExist() bool {
	return !sql.DB.Where(favor).First(favor).RecordNotFound()
}

func (favor *Favor) Delete() error {
	return sql.DB.Delete(favor).Error
}

func (favor *Favor) Insert() error {
	return sql.DB.Create(favor).Error
}

func (favor *Favor) Count() (count int) {
	sql.DB.Model(favor).Where("topic_id = ?", favor.TopicID).Count(&count)
	return
}
