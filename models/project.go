package models

// 项目, 可针对项目创建topic集, 
type Project struct {
	BaseModel
	Name					string			`json:"name" gorm:"not null"`					// 项目名称
	AuthorID			int					`json:"authorID" gorm:"index"`
	Auth					int					`json:"auth" gorm:"tinyint"`					// 权限, 私有, 公开, 部门
	Status				int					`json:"going" gorm:""`								// 项目状态 1: 进行中, 2: 关闭
	Members				[]*User			`json:"members" gorm:"-"`							// 项目成员
}
