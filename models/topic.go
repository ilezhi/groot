package models

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
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
	LikeCount		int						`json:"likeCount" gorm:"-"`
	ComtCount		int						`json:"comtCount" gorm:"-"`
	FavorCount	int						`json:"favorCount" gorm:"-"`
	NickName		string				`json:"nickName" gorm:"-"`
	Avatar			string				`json:"avatar" gorm:"-"`
}

func (topic *Topic) BeforeCreate() error {
	fmt.Println("before create")
	topic.UpdatedAt = time.Now().Unix()
	return nil
}

func (topic *Topic) BeforeUpdate() error {
	fmt.Println("before Update")
	topic.UpdatedAt = time.Now().Unix()
	return nil
}

func (topic *Topic) Validate() error {
	fmt.Println("验证topic", topic)
	return validator.New().Struct(topic)
}

func (topic *Topic) IsExist() bool {
	return !sql.DB.First(topic, topic.ID).RecordNotFound()
}

func (topic *Topic) All() ([]*Topic, error) {
	return topic.SearchByPage("1 = ?", 1)
}

func (topic *Topic) Awesomes() ([]*Topic, error) {
	return topic.SearchByPage("t.awesome = ?", 1)
}

func (topic *Topic) Department(deptID uint) (topics []*Topic, err error) {
	return topic.SearchByPage("u.dept_id = ?", deptID)
}

func (topic *Topic) UnSolved() (topics []*Topic, err error) {
	return topic.SearchByPage("u.id = ? AND t.answer_id = 0", topic.AuthorID)
}

func (topic *Topic) Solved() (topics []*Topic, err error) {
	return topic.SearchByPage("u.id = ? AND t.answerID <> 0", topic.AuthorID)
}

func (topic *Topic) CommentAsAnswer() (topics []*Topic, err error) {
	joins := "INNER comments c ON t.answer_id = c.id"
	err = PageTopics(topic.UpdatedAt).Joins(joins).Where("u.id = ?", topic.AuthorID).Scan(&topics).Error
	if err != nil {
		return
	}

	err = SetTag(&topics)
	return
}

/**
 * id 分类id
 */
func (topic *Topic) GetByCategory(id uint) (topics []*Topic, err error) {
	joins := "JOIN favor f ON f.topic_id = t.id AND f.id = ?"
	err = PageTopics(topic.UpdatedAt).Joins(joins, id).Scan(&topics).Error
	if err != nil {
		return
	}

	err = SetTag(&topics)
	return
}

func (topic *Topic) SharedList() (topics []*Topic, err error) {
	return topic.SearchByPage("t.shared = 1 AND t.author_id = ?", topic.AuthorID)
}

func (topic *Topic) GetByTag(id uint) (topics []*Topic, err error) {
	joins := "JOIN topic_tags tt ON tt.topic_id = t.id AND tt.id = ?"
	err = PageTopics(topic.UpdatedAt).Joins(joins, topic.ID).Where("u.id = ?", topic.AuthorID).Scan(&topics).Error
	return
}

func (topic *Topic) SearchByPage(where string, val interface{}) (topics []*Topic, err error) {
	err = PageTopics(topic.UpdatedAt).Where(where, val).Scan(&topics).Error
	if err != nil {
		return
	}

	err = SetTag(&topics)
	return
}

func (topic *Topic) GetTags() error {
	var tags []*Tag
	fields := "t.id, t.name"
	joins := "INNER JOIN topic_tags tt ON tt.tag_id = t.id AND tt.topic_id = ?"
	err := sql.DB.Table("tags t").Select(fields).Joins(joins, topic.ID).Scan(&tags).Error
	return err
}

func (topic *Topic) Insert(tags *[]uint) error {
	tx := sql.DB.Begin()
	err := tx.Create(topic).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, id := range *tags {
		tt := &TopicTag{
			TopicID: topic.ID,
			TagID: id,
		}
		err = tx.Create(tt).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (topic *Topic) Update(tags *[]uint) error {
	tx := sql.DB.Begin()
	err := tx.Model(topic).Update("content", "shared", "updated_at").Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除原tag
	err = tx.Where("topic_id = ?", topic.ID).Delete(TopicTag{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 添加tag
	for _, id := range *tags {
		tt := &TopicTag{
			TopicID: topic.ID,
			TagID: id,
		}
		err = tx.Create(tt).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func PageTopics(lastID int64) *gorm.DB {
	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	fields := `t.id, t.title, substring(t.content, 1, 140) as content,
						t.view, t.top, t.awesome, t.updated_at, t.answer_id, u.avatar, u.nickname`

	joins := "INNER JOIN users u ON t.author_id = u.id"
	where := "t.updated_at < ? AND t.issue = 1"
	order := "t.updated_at DESC"
	return sql.DB.Table("topics t").Select(fields).Where(where, lastID).Joins(joins).Order(order)
}

func SetTag(topics *[]*Topic) error {
	for _, topic := range *topics {
		err := topic.GetTags()
		if err != nil {
			return err
		}
	}

	return nil
}
