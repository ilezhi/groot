package model

import "time"

type Role struct {
	Id					int				`xorm:"int(3) pk autoincr"`
	Name				string		`xorm:"varchar(20) unique"`
	Description	string		`xorm:"varchar(100)"`
	CreatedAt		time.Time	`xorm:"datetime created"`
	UpdatedAt		time.Time	`xorm:"datetime updated"`
}
