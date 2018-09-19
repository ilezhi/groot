package models

// 点赞, topic, comment, reply
type Good struct {
	BaseModel
	UserID			uint					`json:"userID" gorm:"index"`
	TargetID		uint					`json:"targetID" gorm:"index"`
	Type				string				`json:"type"`
}
