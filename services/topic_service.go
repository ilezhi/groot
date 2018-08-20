package service

import (
	. "groot/model"
)

type TopicService struct {
	GetList() []Topic
	GetByID(id uint) Topic
	Create(topic *Topic) (Topic, error)
	UpdateByID(topic *Topic, id uint) (Topic, error)
	DeleteByID(id uint) (Topic, bool)
	// 保存成草稿
	saveDraft(topic *Topic) bool
	// 置顶
	SetTop(isTop bool) bool
	// 发布, 新增时默认发布, 如果保存
	Issue(issue bool) bool
}

type topicService struct {}

func (tc *topicService) Create() []Topic {

}
