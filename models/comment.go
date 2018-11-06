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
	Nickname	string			`json:"nickname" gorm:"-"`
	Avatar		string			`json:"avatar" gorm:"-"`
	Title			string			`json:"title" gorm:"-"`
	Shared    bool  			`json:"shared" gorm:"-"`
	RID       uint        `json:"rid" gorm:"-"`
}

// TODO: topic
func (comt *Comment) Save(topic *Topic) error {
	now := time.Now().Unix()

	tx := sql.DB.Begin()

	err := tx.Create(comt).Error
	if err != nil {
		tx.Rollback()
		return err	
	}

	topic.ActiveAt = now
	err = tx.Model(topic).Update("active_at").Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (comt *Comment) Update() error {
	return sql.DB.Model(comt).Update("content", "updated_at").Error
}

func (comt *Comment) Delete() error {
	return sql.DB.Delete(comt).Error
}

func (c *Comment) GetReplies() error {
	var replies []*Reply
	fields := `r.id, r.content, r.comment_id, r.author_id, r.receiver_id, r.updated_at, r.created_at,
						 au.name, au.avatar, ru.name as receiver_name, ru.avatar as receiver_avatar`
	joinsUser := "JOIN users au ON r.author_id = au.id"
	joinsReceiver := "JOIN users ru ON r.receiver_id = ru.id"
	order := "r.created_at ASC"

	err := sql.DB.Table("replies r").Select(fields).Where("r.comment_id = ?", c.ID).Joins(joinsUser).Joins(joinsReceiver).Order(order).Scan(&replies).Error
	if err != nil {
		return err
	}

	c.Replies = replies
	return nil
}

func (c *Comment) Count() (count int) {
	comtCount := 0
	sql.DB.Model(&Comment{}).Where("topic_id = ?", c.TopicID).Count(&comtCount)

	reply := new(Reply)
	reply.TopicID = c.TopicID
	replyCount := reply.Count()
	return comtCount + replyCount
}
