package models

type ProjectMember struct {
	BaseModel
	PorjectID     int		`json:"projectID" gorm:"index"`
	MemberID			int		`json:"memberID" gorm:"index"`
}
