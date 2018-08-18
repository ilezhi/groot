package model

type Reply struct {
	BaseModel
	Content			string			`gorm:"type:text"`
	CommentID		uint
	AuthorID		uint
	ReceiverID	uint
	TotalGood		uint				`gorm:"default:'0'"`
	Author			User				`gorm:"-"`
	Receiver		User				`gorm:"-"`
}
