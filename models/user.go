package models

import (
	"time"
	"gopkg.in/go-playground/validator.v9"
	sql "groot/db"
)

type User struct {
	ID				int				`json:"id" gorm:"primary_key;auto_increment"`
	Name			string		`json:"name" gorm:"size:15;not null" validate:"min=2,max=5,required"`
	Nickname	string		`json:"nickname" gorm:"size:50;default:''"`
	Email			string		`json:"email" gorm:"size:50;unique_index" validate:"required,email"`
	Password	string		`json:"password"`
	Gender		int				`json:"gender" gorm:"type:tinyint"`
	Phone			string		`json:"phone" gorm:"type:char(11)"`
	Avatar		string		`json:"avatar" gorm:"type:varchar(2048)"`
	Birthday	string		`json:"birthday" gorm:"type:char(10)"`
	DeptID		int				`json:"deptID"`
	Token			int64			`json:"token"`
	SecretKey	string		`json:"secreKey"`
	IsVerify	bool			`json:"isVerify" gorm:"default:'0'"`		// 默认账号需要邮箱激活验证
	IsLock		bool			`json:"isLock" gorm:"default:'0'"`		// 0: 不锁, 1:锁定
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
	IsAdmin		bool			`json:"isAdmin" gorm:"-"`
}

func (user *User) Validate() error {
	valid := validator.New()
	err := valid.Struct(user)

	return err
}

func (user *User) Save() error {
	return sql.DB.Create(user).Error
}

func (user *User) Find() error {
	return sql.DB.First(user, user.ID).Error
}

func (user *User) FindByEmail() error {
	return sql.DB.Where("email = ?", user.Email).First(user).Error
}
