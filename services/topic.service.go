package service

import (
	"fmt"
	. "groot/db"
	. "groot/model"
)

type ITopic interface {
	GetList() []Topic
	GetByID(id uint) Topic
	Create(topic *Topic) error
	UpdateByID(topic *Topic, id uint) error
	DeleteByID(id uint) (Topic, bool)
	// 保存成草稿
	saveDraft(topic *Topic) bool
	// 置顶
	SetTop(isTop bool) bool
	// 发布, 新增时默认发布, 如果保存
	Issue(issue bool) bool
}

type TopicService struct {}

var TS = TopicService{}

/**
 * 新增话题
 */
func (ts *TopicService) Create(topic *Topic) error {
	fmt.Println("service", topic)
	tx := DB.Begin()

	err := tx.Create(topic).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	
	for _, id := range topic.Tags {
		val := id.(float64)
		tag := &TopicTag{
			TopicID: topic.ID,
			TagID: uint(val),
		}

		err = tx.Create(tag).Error

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
