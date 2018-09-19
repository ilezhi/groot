package services

import (
	"fmt"
	"groot/db"
	. "groot/models"
)

type ITopic interface {
	GetList() ([]*Topic, error)
	GetByID(id uint) (*Topic, error)
	Create(topic *Topic) error
	UpdateByID(topic *Topic, id uint) error
	DeleteByID(id uint) (*Topic, bool)
	// 保存成草稿
	saveDraft(topic *Topic) bool
	// 置顶
	SetTop(isTop bool) bool
	// 发布, 新增时默认发布, 如果保存
	Issue(issue bool) bool
}

type topicService struct {}

var TopicService = topicService{}

/**
 * 获取topic list
 */
func (ts *topicService) GetList(issue bool) ([]*Topic, error) {
	var topics []*Topic
	var err error
	if issue {
		err = db.DB.Where("issue = ?", true).Limit(2).Find(&topics).Error
	} else {
		err = db.DB.Limit(2).Find(&topics).Error
	}

	return topics, err
}

/**
 * 根据id获取topic
 */
func (ts *topicService) GetByID(id uint) (*Topic, error) {
	var topic Topic
	var tags []*Tag

	err := db.DB.First(&topic, id).Error
	if err != nil {
		return nil, err
	}
	
	topic.View++
	topic.UpdateView()
	
	tags, err = TagService.FindByTopicID(id)

	if err != nil {
		return nil, err
	}

	topic.Tags = tags
	fmt.Println("topic", topic)
	return &topic, err
}

/**
 * 新增话题
 */
func (ts *topicService) Create(topic *Topic) error {
	// fmt.Println("service", topic)
	// tx := DB.Begin()

	// err := tx.Create(topic).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	
	// for _, id := range topic.Tags {
	// 	val := id.(float64)
	// 	tag := &TopicTag{
	// 		TopicID: topic.ID,
	// 		TagID: uint(val),
	// 	}

	// 	err = tx.Create(tag).Error

	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// }

	// tx.Commit()
	return nil
}
