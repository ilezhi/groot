package models

import (
	sql "groot/db"
)

type Department struct {
	BaseModel
	Name					string			`json:"name" gorm:"size:30"`
	Description		string			`json:"description"`
}


func (d *Department) List() ([]*Department, error) {
	var depts []*Department
	err := sql.DB.Find(&depts).Error
	return depts, err
}
