package models

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
