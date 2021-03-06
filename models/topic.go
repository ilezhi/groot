package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	sql "groot/db"
)

type TopicParams struct {
	Title 	string		`json:"title"`
	Content string		`json:"content"`
	Tags 		[]int			`json:"tags"`
	Shared 	bool			`json:"shared"`
}

type Topic struct {
	BaseModel
	Title					string				`json:"title" gorm:"type:varchar(100);index;not null" validate:"min=10,max=30,required"`
	Content				string				`json:"content" gorm:"type:text"`
	Shared				bool					`json:"shared" gorm:"default:'0'"`
	AuthorID			int						`json:"authorID" gorm:"index" validate:"required,numeric"`
	View					int						`json:"view" gorm:"default:'0'"`			// 浏览量
	Top						bool					`json:"top" gorm:"default:'0'"`				// 置顶
	Awesome				bool					`json:"awesome" gorm:"default:'0'"`		// 精华
	Issue					bool					`json:"issue" gorm:"default:'1'"`			// 默认发布
	ActiveAt			int64					`json:"activeAt"`
	AnswerID			int						`json:"answerID"`
	Answer				*Comment			`json:"answer" gorm:"-"`
	Tags					[]*Tag				`json:"tags,-" gorm:"-"`
	LikeCount			int						`json:"likeCount" gorm:"-"`
	ComtCount			int						`json:"comtCount" gorm:"-"`
	FavorCount		int						`json:"favorCount" gorm:"-"`
	Nickname			string				`json:"nickname" gorm:"-"`
	Avatar				string				`json:"avatar" gorm:"-"`
	DeptID				int						`json:"deptID" gorm:"-"`
	IsLike				bool					`json:"isLike" gorm:"-"`
	IsFavor				bool					`json:"isFavor" gorm:"-"`
	CategoryID		int						`json:"categoryID" gorm:"-"`
	IsFull				bool					`json:"isFull" gorm:"-"`
	LastNickname 	string				`json:"lastNickname" gorm:"-"`
	LastAvatar  	string				`json:"lastAvatar" gorm:"-"`
}

func (topic *Topic) BeforeCreate() error {
	topic.ActiveAt = time.Now().Unix()
	return nil
}

func (topic *Topic) Validate() error {
	return validator.New().Struct(topic)
}

func (topic *Topic) IsExist() bool {
	return !sql.DB.First(topic, topic.ID).RecordNotFound()
}

func (topic *Topic) TopTopics() ([]*Topic, error) {
	return topic.SearchByPage("t.top = ?", 1, 5)
}

func (topic *Topic) All(size int) ([]*Topic, error) {
	return topic.SearchByPage("1 = ? AND t.top <> 1", 1, size)
}

func (topic *Topic) Awesomes(size int) ([]*Topic, error) {
	return topic.SearchByPage("t.awesome = ?", 1, size)
}

func (topic *Topic) Department(deptID int, size int) (topics []*Topic, err error) {
	return topic.SearchByPage("au.dept_id = ?", deptID, size)
}

func (topic *Topic) UnSolved(size int) (topics []*Topic, err error) {
	return topic.SearchByPage("au.id = ? AND t.answer_id = 0", topic.AuthorID, size)
}

func (topic *Topic) Solved(size int) (topics []*Topic, err error) {
	return topic.SearchByPage("au.id = ? AND t.answer_id <> 0", topic.AuthorID, size)
}

func (topic *Topic) CommentAsAnswer(size int, uid int) (topics []*Topic, err error) {
	joins := "JOIN comments c ON t.answer_id = c.id AND c.author_id = ?"
	err = PageTopics(topic.ActiveAt, size).Joins(joins, uid).Scan(&topics).Error
	if err != nil {
		return
	}

	err = SetTag(&topics)
	return
}

func (topic *Topic) FindByID() error {
	fields := `t.*, u.nickname, u.avatar`
	joins := "INNER JOIN users u ON u.id = t.author_id"
	err := sql.DB.Table("topics t").Select(fields).Where("t.id = ?", topic.ID).Joins(joins).Scan(topic).Error
	if err != nil {
		return err
	}

	err = topic.GetTags()
	return err
}

func (topic *Topic) FindFullByID() error {
	fields := `t.id, t.title, substring(t.content, 1, 140) as content, t.author_id,
						t.view, t.top, t.shared, t.awesome, t.active_at, t.created_at, t.answer_id, au.avatar, au.nickname, au.dept_id,
						lu.nickname as last_nickname, lu.avatar as last_avatar`
	
	joins := "JOIN users au ON t.author_id = au.id"
	lastPost := `LEFT JOIN (
		SELECT d.* from (
			SELECT author_id, topic_id, updated_at from comments
			UNION
			SELECT author_id, topic_id, updated_at from replies
			ORDER BY updated_at DESC
		) d GROUP BY d.topic_id
	) lp on lp.topic_id = t.id`
	lastJoins := "left join users lu on lu.id = lp.author_id"
	where := "t.id = ?"
	err := sql.DB.Table("topics t").Select(fields).Where(where, topic.ID).Joins(joins).Joins(lastPost).Joins(lastJoins).Scan(topic).Error
	if err != nil {
		return err
	}

	err = topic.GetTags()
	return err
}

func (topic *Topic) GetCount(uid int) {
	like := new(Like)
	like.TargetID = topic.ID
	like.Type = "topic"
	like.UserID = uid

	favor := new(Favor)
	favor.TopicID = topic.ID
	favor.UserID = uid

	topic.LikeCount = like.Count()
	topic.FavorCount = favor.Count()
	topic.IsLike = like.IsExist()
	topic.IsFavor = favor.IsExist()
	topic.CategoryID = favor.CategoryID
	topic.GetComtCount()
}

/**
 * 获取评论总数, 包含回复数量
 */
func (topic *Topic) GetComtCount() {
	comt := new(Comment)
	comt.TopicID = topic.ID
	topic.ComtCount = comt.Count()
}

/**
 * id 分类id
 */
func (topic *Topic) GetByCategory(id int, size int) (topics []*Topic, err error) {
	joins := "JOIN favors f ON f.topic_id = t.id AND f.category_id = ?"
	err = PageTopics(topic.ActiveAt, size).Joins(joins, id).Scan(&topics).Error
	if err != nil {
		return
	}

	for _, t := range topics {
		t.GetComtCount()
	}

	err = SetTag(&topics)
	return
}

func (topic *Topic) SharedList(size int) (topics []*Topic, err error) {
	return topic.SearchByPage("t.shared = 1 AND t.author_id = ?", topic.AuthorID, size)
}

func (topic *Topic) GetByTag(id int, size int) (topics []*Topic, err error) {
	joins := "JOIN topic_tags tt ON tt.topic_id = t.id AND tt.tag_id = ?"
	err = PageTopics(topic.ActiveAt, size).Joins(joins, id).Where("au.id = ?", topic.AuthorID).Scan(&topics).Error
	if err != nil {
		return
	}

	for _, t := range topics {
		t.GetComtCount()
	}

	err = SetTag(&topics)
	return
}

func (topic *Topic) SearchByPage(where string, val interface{}, size int) (topics []*Topic, err error) {
	err = PageTopics(topic.ActiveAt, size).Where(where, val).Scan(&topics).Error
	if err != nil {
		return
	}

	for _, t := range topics {
		t.GetComtCount()
	}

	err = SetTag(&topics)
	return
}

func (topic *Topic) GetTags() error {
	var tags []*Tag
	fields := "t.id, t.name, t.created_at, t.author_id"
	joins := "INNER JOIN topic_tags tt ON tt.tag_id = t.id AND tt.topic_id = ?"

	err := sql.DB.Table("tags t").Select(fields).Joins(joins, topic.ID).Scan(&tags).Error
	topic.Tags = tags

	return err
}

func (topic *Topic) Save(tags *[]int) error {
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

func (topic *Topic) Update(tags *[]int) error {
	topic.ActiveAt = time.Now().Unix()
	tx := sql.DB.Begin()
	err := tx.Model(topic).Select("content", "shared", "active_at").Updates(*topic).Error
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

func (topic *Topic) UpdateView() error {
	return sql.DB.Model(topic).UpdateColumn("view", topic.View).Error
}

func (topic *Topic) UpdateField(field string, value interface{}) error {
	now := time.Now().Unix()
	fields := make(map[string]interface{})
	fields[field] = value
	fields["active_at"] = now
	return sql.DB.Model(topic).UpdateColumns(fields).Error
}

func (topic *Topic) GetComments(uid int) (comments []*Comment, err error) {
	fields := `c.id, c.content, c.topic_id, c.updated_at, c.author_id, c.created_at,
							u.nickname, u.avatar`

	joins := "JOIN users u ON c.author_id = u.id"
	order := "c.created_at ASC"

	err = sql.DB.Table("comments c").Select(fields).Joins(joins).Where("c.topic_id = ?", topic.ID).Order(order).Scan(&comments).Error
	if err != nil {
		return
	}

	// 获取评论回复
	for _, comt := range comments {
		like := new(Like)
		like.TargetID = comt.ID
		like.Type = "comment"
		like.UserID = uid
		comt.LikeCount = like.Count()
		comt.IsLike = like.IsExist()

		err = comt.GetReplies(uid)
		if err != nil {
			return
		}
	}

	return
}

func PageTopics(lastID int64, size int) *gorm.DB {
	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	fields := `t.id, t.title, substring(t.content, 1, 140) as content, t.author_id,
						t.view, t.top, t.shared, t.awesome, t.active_at, t.created_at, t.answer_id, au.avatar, au.nickname, au.dept_id,
						lu.nickname as last_nickname, lu.avatar as last_avatar`
	
	joins := "JOIN users au ON t.author_id = au.id"
	lastPost := `LEFT JOIN (
		SELECT d.* from (
			SELECT author_id, topic_id, updated_at from comments
			UNION
			SELECT author_id, topic_id, updated_at from replies
			ORDER BY updated_at DESC
		) d GROUP BY d.topic_id
	) lp ON lp.topic_id = t.id`
	lastJoins := "LEFT JOIN users lu ON lu.id = lp.author_id"
	where := "t.active_at < ? AND t.issue = 1"
	order := "t.active_at DESC"
	return sql.DB.Table("topics t").Select(fields).Where(where, lastID).Limit(size).Joins(joins).Joins(lastPost).Joins(lastJoins).Order(order)
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
