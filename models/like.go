package models

import (
	sql "groot/db"
)

// 点赞, topic, comment, reply
type Like struct {
	BaseModel
	UserID			uint					`json:"userID" gorm:"index"`
	TargetID		uint					`json:"targetID" gorm:"index"`
	Type				string				`json:"type"`
}

func (like *Like) IsExist() bool {
	return !sql.DB.Where(like).Find(like).RecordNotFound()
}

func (like *Like) Save() error {
	return sql.DB.Create(like).Error
}

func (like *Like) Delete() error {
	return sql.DB.Delete(like).Error
}
