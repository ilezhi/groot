package model

type Topic struct {
	BaseModel
	Title				string			`gorm:"type:varchar(100);index;not null"`
	Content			string			`gorm:"type:text"`
	Tags				[]*Tag			`gorm:"-"`
	Comments		[]*Comment	`gorm:"-"`
	AuthorID		uint
	View				uint				`gorm:"default:'0'"`			// 浏览量
	TotalComt		uint				`gorm:"default:'0'"`			// 评论和回复总数
	TotalGood		uint				`gorm:"default:'0'"`			// 赞数
	top					bool				`gorm:"default:'0'"`			// 置顶
	issue				bool				`gorm:"default:'1'"`			// 默认发布
	NickName		string			`gorm:"-"`
	Avatar			string			`gorm:"-"`
}
