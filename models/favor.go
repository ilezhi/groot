package models

// 收藏
type Favor struct {
	BaseModel
	UserID			uint					`json:"userID" gorm:"index"`
	TopicID			uint					`json:"topicID" gorm:"index"`
	CategoryID	uint					`json:"categoryID"`
}
