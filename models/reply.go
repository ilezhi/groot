package models

import (
	"time"
	sql "groot/db"
)

type Reply struct {
	BaseModel
	Content			string			`json:"content" gorm:"type:text" validate:"required"`
	CommentID		uint				`json:"commentID" validate:"required,numeric"`
	AuthorID		uint				`json:"authorID"`
	ReceiverID	uint				`json:"receiverID" validate:"required,numeric"`
	UpdatedAt		int64				`json:"updatedAt"`
	Author			*User				`json:"author" gorm:"-"`
	Receiver		*User				`json:"receiver" gorm:"-"`
}

func (reply *Reply) BeforeCreate() error {
	reply.UpdatedAt = time.Now().Unix()
	return nil
}

func (reply *Reply) Save() error {
	return sql.DB.Create(reply).Error
}
