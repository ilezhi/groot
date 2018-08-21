package model

type Topic struct {
	BaseModel
	Title				string			`json:"title" gorm:"type:varchar(100);index;not null"`
	Content			string			`json:"content" gorm:"type:text"`
	Tags				[]*Tag			`json:"tags" gorm:"-"`
	Comments		[]*Comment	`json:"comments" gorm:"-"`
	AnwserID		uint				`json:"anwserID"`			
	AuthorID		uint				`json:"authorID" gorm:"index"`
	ProjectID		uint				`json:"projectID"`
	View				uint				`json:"view" gorm:"default:'0'"`			// 浏览量
	TotalComt		uint				`json:"totalComt" gorm:"default:'0'"`			// 评论和回复总数
	TotalGood		uint				`json:"totalGood" gorm:"default:'0'"`			// 赞数
	TotalFavor  uint				`json:"totalFavor" gorm:"default:'0'"`			// 收藏数
	Top					bool				`json:"top" gorm:"default:'0'"`			// 置顶
	Issue				bool				`json:"issue" gorm:"default:'1'"`			// 默认发布
	NickName		string			`json:"nickName" gorm:"-"`
	Avatar			string			`json:"avatar" gorm:"-"`
}
