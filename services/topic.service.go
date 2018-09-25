package services

import (
	"time"
	sql "groot/db"
	. "groot/models"
)

type ITopic interface {
	Find(lastID int64, size int) ([]*Topic, error)
	FindAwesome(lastID int64, size int) ([]*Topic, error)
	FindByID(id uint) (*Topic, error)
	FindListByQuery(lastID int64, size int, query map[string]interface{})(map[string]interface{}, error)
	ByID(id uint) (*Topic, error)
	Create(topic *Topic, tags *[]uint) error
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
}

type topicService struct {}

var TopicService = topicService{}

/**
 * 获取topic list
 */
func (ts *topicService) Find(lastID int64, size int) ([]*Topic, error) {
	var topics []*Topic

	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	err := sql.DB.Order("updated_at desc").Where("updated_at < ? AND issue = ?", lastID, true).Limit(size).Find(&topics).Error
	return topics, err
}

/**
 *
 */
func (ts *topicService) FindAwesome(lastID int64, size int) ([]*Topic, error) {
	var topics []*Topic

	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	err := sql.DB.Order("updated_at desc").Where("updated_at < ? AND issue = 1 AND awesome = 1", lastID).Limit(size).Find(&topics).Error
	return topics, err
}

/**
 * 根据id获取topic, 同时获取tags
 */
func (ts *topicService) FindByID(id uint) (*Topic, error) {
	var tags []*Tag

	topic, err := ts.ByID(id)
	if err != nil {
		return nil, err
	}

	tags, err = TagService.FindByTopicID(id)
	topic.Tags = tags

	return topic, err
}

func (ts *topicService) FindListByQuery(lastID int64, size int, query map[string]interface{}) (map[string]interface{}, error) {
	var topics []*Topic

	if lastID == -1 {
		lastID = time.Now().Unix()
	}

	err := sql.DB.Order("updated_at desc").Where("updated_at < ? AND issue = 1", lastID).Where(query).Limit(size).Find(&topics).Error

	if err != nil {
		return nil, err
	}

	// 获取tag
	err = TagService.FindByTopics(topics)
	count := ts.Count(query)

	data := map[string]interface{}{
		"total": count,
		"list": topics,
	}

	return data, err
}

func (ts *topicService) ByID(id uint) (*Topic, error) {
	var topic Topic

	err := sql.DB.First(&topic, id).Error

	return &topic, err	
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

func (ts *topicService) FindAndUpdate(id uint, content string , tags *[]uint) (*Topic, error) {
	tx := sql.DB.Begin()

	var topic Topic
	err := sql.DB.First(&topic, id).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	topic.Content = content
	topic.UpdatedAt = time.Now().Unix()

	err =	sql.DB.Save(&topic).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = TagService.DeleteByTopicID(id)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = ts.saveTag(id, tags)
	
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	topic.Tags, err = TagService.FindByTopicID(id)

	if err != nil {
		return nil, err
	}

	return &topic, nil
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
