package model

type Comment struct {
	BaseModel
	Content		string			`json:"content" gorm:"type:text"`
	Replies		[]*Reply		`json:"replies" gorm:"-"`
	AuthorID	uint				`json:"authorID"`
	TopicID		uint				`json:"topicID"`
	TotalGood	uint				`json:"totalGood" gorm:"default:'0'"`
	NickName	string			`json:"nickName" gorm:"-"`
	Avatar		string			`json:"avatar" gorm:"-"`
}
