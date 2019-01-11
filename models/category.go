package models

import (
	sql "groot/db"
)

// 类别: 将收藏的topic 进行分类
type Category struct {
	BaseModel
	Name				string			`json:"name" gorm:"size:30;not null"`
	UserID			int					`json:"userID"`
	Count				int					`json:"count" gorm:"-"`
}

func (c *Category) IsExist() bool {
	return !sql.DB.Where(c).Find(c).RecordNotFound()
}

func (c *Category) Save() error {
	return sql.DB.Create(c).Error
}

func (c *Category) GroupBy() ([]*Category, error) {
	var categories []*Category
	fields := "c.name, IFNULL(f.count, 0) as count, c.user_id, c.id, c.created_at"
	joins := `LEFT JOIN (
		select category_id, count(*) as count from favors group by category_id
		) f ON c.id = f.category_id`

	err := sql.DB.Table("categories c").Select(fields).Where("c.user_id = ?", c.UserID).Joins(joins).Scan(&categories).Error
	return categories, err
}
