package models

import (
	"fmt"
	"time"
	sql "groot/db"
)

type Reply struct {
	BaseModel
	Content					string			`json:"content" gorm:"type:text" validate:"required"`
	CommentID				uint				`json:"commentID" validate:"required,numeric"`
	AuthorID				uint				`json:"authorID"`
	ReceiverID			uint				`json:"receiverID" validate:"required,numeric"`
	UpdatedAt				int64				`json:"updatedAt"`
	Name						string			`json:"name" gorm:"-"`
	Avatar					string			`json:"avatar" gorm:"-"`
	ReceiverName		string			`json:"receiverName" gorm:"-"`
	ReceiverAvatar	string			`json:"receiverAvatar" gorm:"-"`
}

func (reply *Reply) BeforeCreate() error {
	fmt.Println("before create")
	reply.UpdatedAt = time.Now().Unix()
	return nil
}

func (reply *Reply) BeforeUpdate() error {
	fmt.Println("before Update")
	reply.UpdatedAt = time.Now().Unix()
	return nil
}

func (reply *Reply) Save() error {
	return sql.DB.Create(reply).Error
}
