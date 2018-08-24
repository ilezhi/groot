package model

import "fmt"
import "gopkg.in/go-playground/validator.v9"
import . "groot/db"

type Topic struct {
	BaseModel
	Title				string				`json:"title" gorm:"type:varchar(100);index;not null" validate:"min=2,max=30,required"`
	Content			string				`json:"content" gorm:"type:text"`
	Tags				[]interface{}	`json:"tags,-" gorm:"-" validate:"gte=1,dive,numeric"`
	Comments		[]interface{}	`json:"comments,-" gorm:"-"`
	AnwserID		uint					`json:"anwserID"`
	AuthorID		uint					`json:"authorID" gorm:"index" validate:"required,numeric"`
	ProjectID		uint					`json:"projectID" validate:"numeric"`
	View				uint					`json:"view" gorm:"default:'0'"`			// 浏览量
	TotalComt		uint					`json:"totalComt" gorm:"default:'0'"`			// 评论和回复总数
	TotalGood		uint					`json:"totalGood" gorm:"default:'0'"`			// 赞数
	TotalFavor  uint					`json:"totalFavor" gorm:"default:'0'"`			// 收藏数
	Top					bool					`json:"top" gorm:"default:'0'"`			// 置顶
	Issue				bool					`json:"issue" gorm:"default:'1'"`			// 默认发布
	NickName		string				`json:"nickName" gorm:"-"`
	Avatar			string				`json:"avatar" gorm:"-"`
}

func (topic *Topic) Validate() error {
	fmt.Println("验证topic", topic)
	return validator.New().Struct(topic)
}

func (topic *Topic) UpdateView() error {
	return DB.Model(topic).Update("view", topic.View).Error
}
