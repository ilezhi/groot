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
	Nickname	string			`json:"nickName" gorm:"-"`
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

func (c *Comment) GetReplies() error {
	var replies []*Reply
	fields := `r.content, r.comment_id, r.author_id, r.receiver_id
						 au.name, au.avatar, ru.name as receiver_name, ru.avatar as receiver_avatar`
	joinsUser := "JOIN users au ON r.author_id = au.id"
	joinsReceiver := "JOIN users ru ON r.receiver_id = ru.id"
	order := "r.created_at ASC"

	err := sql.DB.Table("replies r").Select(fields).Where("comment_id = ?", c.ID).Joins(joinsUser).Joins(joinsReceiver).Order(order).Scan(&replies).Error
	if err != nil {
		return err
	}
	c.Replies = replies
	return nil
}
