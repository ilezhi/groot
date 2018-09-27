package services

import (
	"time"
	sql "groot/db"
	. "groot/models"
)

type ITopic interface {
	Find(lastID int64) ([]*Topic, error)
	FindByID(id uint) (*Topic, error)
	FindByQuery(lastID int64, query map[string]interface{}) ([]*Topic, error)
	ByID(id uint) (*Topic, error)
	FindQuestion(userID uint, lastID int64) ([]*Topic, error)
	FindAnswer(userID uint, lastID int64) ([]*Topic, error)
	Create(topic *Topic, tags *[]uint) error
	Update(topic *Topic, params *TopicParams) error
	FindAndUpdate(id uint, content string, tags *[]uint) (*Topic, error)
	DeleteByID(id uint) error
	// 保存成草稿
	saveDraft(topic *Topic) bool
	// 置顶
	SetTop(isTop bool) bool
	// 发布, 新增时默认发布, 如果保存
	Issue(issue bool) bool

	saveTag(id uint, tags *[]uint) error
	FindAndUpdateColumns(id uint, columns interface{}) (*Topic, error)
	Count(query map[string]interface{}) int

	// 收藏
	Favor(favor *Favor) (bool, error)
}

type topicService struct {
	size int
	fields string
}

var TopicService = topicService{size: 50}

/**
 * 获取topic list
 */
func (ts *topicService) Find(lastID int64) ([]*Topic, error) {
	var topics []*Topic

	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	err := sql.DB.Order("updated_at desc").Where("updated_at < ? AND issue = ?", lastID, true).Limit(ts.size).Find(&topics).Error
	return topics, err
}


/**
 * 根据id获取topic, 同时获取tags
 */
func (ts *topicService) FindByID(id uint) (*Topic, error) {
	topic, err := ts.ByID(id)
	if err != nil {
		return nil, err
	}

	tags, err := TagService.FindByTopicID(id)
	topic.Tags = tags

	return topic, err
}

func (ts *topicService) FindByQuery(lastID int64, query map[string]interface{}) ([]*Topic, error) {
	var topics []*Topic

	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	fields := `id, title, substring(content, 1, 140) as content,
						view, total_comt, top, awesome, updated_at`
	err := sql.DB.Select(fields).Order("updated_at desc").Where("updated_at < ? AND issue = 1", lastID).Where(query).Limit(ts.size).Find(&topics).Error

	if err != nil {
		return nil, err
	}

	// 获取tag
	err = TagService.FindByTopics(topics)
	return topics, err
}

func (ts *topicService) ByID(id uint) (*Topic, error) {
	var topic Topic

	err := sql.DB.First(&topic, id).Error

	return &topic, err	
}

/**
 * 我的提问(已解决)
 */ 
func (ts *topicService) FindQuestion(userID uint, lastID int64) ([]*Topic, error) {
	var topics []*Topic

	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	err := sql.DB.Table("topics").Order("updated_at desc").Where("issue = 1 AND answer_id != 0 AND author_id = ?", userID).Where("updated_at < ?", lastID).Limit(ts.size).Scan(&topics).Error
	if err != nil {
		return nil, err
	}

	err = TagService.FindByTopics(topics)
	return topics, err
}

/**
 * 我的回答被设置为答案
 */
func (ts *topicService) FindAnswer(userID uint, lastID int64) ([]*Topic, error) {
	var topics []*Topic

	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	err := sql.DB.Table("topics t").Select("t.*").Joins("INNER JOIN comments c ON t.answer_id = c.id AND c.author_id = ?", userID).Where("t.issue = 1 AND t.updated_at < ?", lastID).Order("t.updated_at desc").Limit(ts.size).Scan(&topics).Error
	return topics, err
}

/**
 * 新增话题
 */
func (ts *topicService) Create(topic *Topic, tags *[]uint) error {
	// fmt.Println("service", topic)
	topic.UpdatedAt = time.Now().Unix()
	tx := sql.DB.Begin()

	err := tx.Create(topic).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = ts.saveTag(topic.ID, tags)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

/**
 * 更新话题
 * 更新content和tags
 */
func (ts *topicService) Update(topic *Topic, params *TopicParams) error {
	topic.Content = params.Content
	id := topic.ID

	tx := sql.DB.Begin()

	err := topic.Update()
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除原tag
	err = TagService.DeleteByTopicID(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 添加新tag
	err = ts.saveTag(id, &params.Tags)
	if err != nil {
		tx.Rollback()
		return err
	}

	topic.Tags, err = TagService.FindByTopicID(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ts *topicService) DeleteByID(id uint) error {
	return sql.DB.Where("id = ?", id).Delete(Topic{}).Error
}

func (ts *topicService) saveTag(id uint, tags *[]uint) error {
	for _, tid := range *tags {
		tag := &TopicTag{
			TopicID: id,
			TagID: tid,
		}

		err := sql.DB.Create(tag).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (ts *topicService) FindAndUpdateColumns(id uint, columns interface{}) (*Topic, error) {
	topic, err := ts.ByID(id)

	if err != nil {
		return nil, err
	}

	err = sql.DB.Model(topic).UpdateColumns(columns).Error

	return topic, err
}

func (ts *topicService) Count(query map[string]interface{}) int {
	var count int
	sql.DB.Model(&Topic{}).Where(query).Count(&count)
	return count
}

func (ts *topicService) Favor(favor *Favor) (bool, error) {
	var exist Favor

	err := sql.DB.Where(favor).Find(&exist).Error
	if err != nil {
		// 收藏
		err = sql.DB.Create(favor).Error
		return true, err
	}

	// 取消收藏
	err = sql.DB.Delete(favor).Error
	return false, err
}
