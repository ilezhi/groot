package services

import (
	"groot/db"
	. "groot/models"
)

type ITopic interface {
	Find(lastID uint, size int) ([]*Topic, error)
	FindByID(id uint) (*Topic, error)
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
}

type topicService struct {}

var TopicService = topicService{}

/**
 * 获取topic list
 */
func (ts *topicService) Find(lastID uint, size int) ([]*Topic, error) {
	var topics []*Topic
	err := db.DB.Order("id desc").Where("id >= ? AND issue = ?", lastID, true).Limit(size).Find(&topics).Error

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

func (ts *topicService) ByID(id uint) (*Topic, error) {
	var topic Topic

	err := db.DB.First(&topic, id).Error

	return &topic, err	
}

/**
 * 新增话题
 */
func (ts *topicService) Create(topic *Topic, tags *[]uint) error {
	// fmt.Println("service", topic)
	tx := db.DB.Begin()

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
	tx := db.DB.Begin()

	var topic Topic
	err := db.DB.First(&topic, id).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	topic.Content = content

	err =	db.DB.Save(&topic).Error

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
	return db.DB.Where("id = ?", id).Delete(Topic{}).Error
}

func (ts *topicService) saveTag(id uint, tags *[]uint) error {
	for _, tid := range *tags {
		tag := &TopicTag{
			TopicID: id,
			TagID: tid,
		}

		err := db.DB.Create(tag).Error

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

	err = db.DB.Model(topic).UpdateColumns(columns).Error

	return topic, err
}
