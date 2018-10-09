package models

import (
	"time"
	sql "groot/db"
)

type Reply struct {
	BaseModel
	Content					string			`json:"content" gorm:"type:text" validate:"required"`
	CommentID				uint				`json:"commentID" validate:"required,numeric"`
	AuthorID				uint				`json:"authorID"`
	TopicID					uint				`json:"topicID"`
	ReceiverID			uint				`json:"receiverID" validate:"required,numeric"`
	UpdatedAt				int64				`json:"updatedAt"`
	Name						string			`json:"name" gorm:"-"`
	Avatar					string			`json:"avatar" gorm:"-"`
	ReceiverName		string			`json:"receiverName" gorm:"-"`
	ReceiverAvatar	string			`json:"receiverAvatar" gorm:"-"`
}

func (reply *Reply) Save(topic *Topic) error {
	now := time.Now().Unix()
	reply.UpdatedAt = now

	tx := sql.DB.Begin()

	err := tx.Create(reply).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	topic.UpdatedAt = now
	err = tx.Model(topic).Update("updated_at").Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (reply *Reply) Count() (count int) {
	sql.DB.Model(reply).Where("topic_id = ?", reply.TopicID).Count(&count)
	return
}
