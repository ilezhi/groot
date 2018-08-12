package model

import "time"

type User struct {
	Id				int64				`xorm:"pk autoincr default 10000"`
	Name			string			`xorm:"varchar(15) not null"`
	Nickname	string			`xorm:"varchar(50) unique index"`
	Email			string			`xorm:"varchar(50) index default ''"`
	Password	string			`xorm:"char(32)"`
	Gender		byte				`xorm:"tinyint(1) default 1"`
	Phone			string			`xorm:"char(11)"`
	Avatar		string			`xorm:"varchar(2048)"`
	Birthday	string			`xorm:"date"`
	Role			Role				`xorm:"role_id int(3) default 2"`
	Lock			bool				`xorm:"bool default false"`
	CreatedAt	time.Time		`xorm:"datetime created"`
	UpdatedAt	time.Time		`xorm:"datetime updated"`
}
