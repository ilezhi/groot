package model

type ProjectMember struct {
	BaseModel
	PorjectID     uint		`json:"projectID" gorm:"index"`
	MemberID			uint		`json:"memberID" gorm:"index"`
}
