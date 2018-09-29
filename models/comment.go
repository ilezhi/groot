package models

import (
	"time"
	sql "groot/db"
)

type Comment struct {
	BaseModel
	Content		string			`json:"content" gorm:"type:text" validate:"required"`
	Replies		[]*Reply		`json:"replies" gorm:"-"`
	AuthorID	uint				`json:"authorID"`
	TopicID		uint				`json:"topicID" validate:"required"`
	UpdatedAt int64				`json:"updatedAt"`
	// TotalGood	int					`json:"totalGood" gorm:"default:'0'"`
	NickName	string			`json:"nickName" gorm:"-"`
	Avatar		string			`json:"avatar" gorm:"-"`
}

func (comt *Comment) Save() error {
	comt.UpdatedAt = time.Now().Unix()
	return sql.DB.Create(comt).Error
}

func (comt *Comment) Update() error {
	comt.UpdatedAt = time.Now().Unix()
	return sql.DB.Model(comt).Update("content", "updated_at").Error
}

func (comt *Comment) Delete() error {
	return sql.DB.Delete(comt).Error
}
