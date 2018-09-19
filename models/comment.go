package models

type Comment struct {
	BaseModel
	Content		string			`json:"content" gorm:"type:text" validate:"required"`
	Replies		[]*Reply		`json:"replies" gorm:"-"`
	AuthorID	uint				`json:"authorID"`
	TopicID		uint				`json:"topicID" validate:"required"`
	TotalGood	int					`json:"totalGood" gorm:"default:'0'"`
	NickName	string			`json:"nickName" gorm:"-"`
	Avatar		string			`json:"avatar" gorm:"-"`
}
