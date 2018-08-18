package model

type Comment struct {
	BaseModel
	Content		string			`gorm:"type:text"`
	Replies		[]*Reply		`gorm:"-"`
	AuthorID	uint
	TopicID		uint
	TotalGood	uint				`gorm:"default:'0'"`
	NickName	string			`gorm:"-"`
	Avatar		string			`gorm:"-"`
}
