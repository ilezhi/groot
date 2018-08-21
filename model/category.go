package model

// 类别: 将收藏的topic 进行分类
type Category struct {
	BaseModel
	Name				string			`json:"name" gorm:"size:30;not null"`
	UserID			uint				`json:"userID"`
}
