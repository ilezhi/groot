package model

type User struct {
	BaseModel
	Name			string		`gorm:"size:15;not null"`
	Nickname	string		`gorm:"size:50;default:''"`
	Email			string		`gorm:"size:50;unique_index"`
	Password	string		`gorm:"type:char(32)"`
	Gender		int				`gorm:"type:tinyint"`
	Phone			string		`gorm:"type:char(11)"`
	Avatar		string		`gorm:"type:varchar(2048)"`
	Birthday	string		`gorm:"type:char(10)"`
	SecretKey	string
	IsAdmin		bool			`gorm:"default:'0'"`
	IsVerify	bool			`gorm:"default:'0'"`		// 默认账号需要邮箱激活验证
	IsLock		bool			`gorm:"default:'0'"`		// 0: 不锁, 1:锁定
}
