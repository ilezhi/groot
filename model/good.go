package model

type Good struct {
	BaseModel
	UserID			uint					`gorm:"index"`
	TargetID		uint					`gorm:"index"`
	Type				string
}
