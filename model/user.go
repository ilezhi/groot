package model

type User struct {
	BaseModel
	Name			string		`json:"name" gorm:"size:15;not null"`
	Nickname	string		`json:"nickName" gorm:"size:50;default:''"`
	Email			string		`json:"email" gorm:"size:50;unique_index"`
	Password	string		`json:"password" gorm:"type:char(32)"`
	Gender		int				`json:"gender" gorm:"type:tinyint"`
	Phone			string		`json:"phone" gorm:"type:char(11)"`
	Avatar		string		`json:"avatar" gorm:"type:varchar(2048)"`
	Birthday	string		`json:"birthday" gorm:"type:char(10)"`
	DeptID		uint			`json:"deptID"`
	Token			int64			`json:"token"`		
	SecretKey	string
	IsAdmin		bool			`json:"isAdmin" gorm:"default:'0'"`
	IsVerify	bool			`json:"isVerify" gorm:"default:'0'"`		// 默认账号需要邮箱激活验证
	IsLock		bool			`json:"isLock" gorm:"default:'0'"`		// 0: 不锁, 1:锁定
}
