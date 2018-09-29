package models

import (
	"fmt"
	"time"
	"gopkg.in/go-playground/validator.v9"
	sql "groot/db"
)

type TopicParams struct {
	Title 	string		`json:"title"`
	Content string		`json:"content"`
	Tags 		[]uint		`json:"tags"`
	Shared 	bool			`json:"shared"`
}

type Topic struct {
	BaseModel
	Title				string				`json:"title" gorm:"type:varchar(100);index;not null" validate:"min=10,max=30,required"`
	Content			string				`json:"content" gorm:"type:text"`
	Shared			bool					`json:"shared" gorm:"default:'0'"`
	AuthorID		uint					`json:"authorID" gorm:"index" validate:"required,numeric"`
	View				uint					`json:"view" gorm:"default:'0'"`			// 浏览量
	Top					bool					`json:"top" gorm:"default:'0'"`				// 置顶
	Awesome			bool					`json:"awesome" gorm:"default:'0'"`		// 精华
	Issue				bool					`json:"issue" gorm:"default:'1'"`			// 默认发布
	UpdatedAt		int64					`json:"updatedAt"`										// 时间戳, 用于排序, 采用lastID排序
	AnswerID		uint					`json:"answerID"`
	Answer			*Comment			`json:"answer" gorm:"-"`
	Tags				[]*Tag				`json:"tags,-" gorm:"-"`
	likeCount		int						`json:"likeCount" gorm:"-"`
	ComtCount		int						`json:"comtCount" gorm:"-"`
	FavorCount	int						`json:"favorCount" gorm:"-"`
	NickName		string				`json:"nickName" gorm:"-"`
	Avatar			string				`json:"avatar" gorm:"-"`
}

func (topic *Topic) Validate() error {
	fmt.Println("验证topic", topic)
	return validator.New().Struct(topic)
}

func (topic *Topic) UpdateView() error {
	return sql.DB.Model(topic).Update("view", topic.View).Error
}

func (topic *Topic) Update() error {
	now := time.Now().Unix()
	topic.UpdatedAt = now

	return sql.DB.Save(topic).Error
}

func (topic *Topic) IsExist(id uint) bool {
	return !sql.DB.First(topic, id).RecordNotFound()
}
