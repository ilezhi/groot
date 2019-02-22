package models

import (
	"time"
	sql "groot/db"
)

type Reply struct {
	BaseModel
	Content					string			`json:"content" gorm:"type:text" validate:"required"`
	CommentID				int					`json:"commentID" validate:"required,numeric"`
	AuthorID				int					`json:"authorID"`
	TopicID					int					`json:"topicID"`
	ReceiverID			int					`json:"receiverID" validate:"required,numeric"`
	Nickname				string			`json:"nickname" gorm:"-"`
	Avatar					string			`json:"avatar" gorm:"-"`
	ReceiverName		string			`json:"receiverName" gorm:"-"`
	ReceiverAvatar	string			`json:"receiverAvatar" gorm:"-"`
	Title						string			`json:"title" gorm:"-"`
	Shared					bool  			`json:"shared" gorm:"-"`
	RID             int        	`json:"rid" gorm:"-"`
	IsLike					bool				`json:"isLike" gorm:"-"`
	LikeCount				int					`json:"likeCount" gorm:"-"`
	Top     				bool				`json:"top" gorm:"-"`  
}

// TODO: topic
func (reply *Reply) Save(topic *Topic) error {
	now := time.Now().Unix()

	tx := sql.DB.Begin()

	err := tx.Create(reply).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(topic).Update("active_at", now).Error
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

func (reply *Reply) ByID() error {
	fields := `r.id, r.content, r.comment_id, r.topic_id, r.author_id, r.receiver_id,
						 r.updated_at, au.nickname, au.avatar,
						 ru.nickname as receiver_name, ru.avatar as receiver_avatar`
	joinsAU := `JOIN users au ON r.author_id = au.id`
	joinsRU	:= `JOIN users ru ON r.receiver_id = ru.id`

	return sql.DB.Table("replies r").Select(fields).Where("r.id = ?", reply.ID).Joins(joinsAU).Joins(joinsRU).Scan(reply).Error
}
