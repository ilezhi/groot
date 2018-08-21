package model

type Reply struct {
	BaseModel
	Content			string			`json:"content" gorm:"type:text"`
	CommentID		uint				`json:"commentID"`
	AuthorID		uint				`json:"authorID"`
	ReceiverID	uint				`json:"receiverID"`
	TotalGood		uint				`json:"totalGood" gorm:"default:'0'"`
	Author			User				`json:"author" gorm:"-"`
	Receiver		User				`json:"receiver" gorm:"-"`
}
