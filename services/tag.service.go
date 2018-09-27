package services

import (
	sql "groot/db"
	. "groot/models"
)

type ITag interface {
	Find() ([]*Tag, error)
	FindByID(id uint) (*Tag, error)
	FindByIDs(ids []uint) ([]*Tag, error)
	FindByTopicID(id uint) (*Tag, error)
	FindByName(name string) ([]*Tag, error)
	FindOneByName(name string) (*Tag, error)
	FindByTopics(topics []*Topic)(error)
	Create(tag *Tag) error
	DeleteByTopicID(id uint) error
}

type tagService struct{}

var TagService = tagService{}

func (ts *tagService) Find() ([]*Tag, error) {
	var tags []*Tag
	err := sql.DB.Find(&tags).Error

	return tags, err
}

func (ts *tagService) FindByID(id uint) (*Tag, error) {
	var tag Tag

	err := sql.DB.Find(&tag, id).Error

	return &tag, err
}

func (ts *tagService) FindByIDs(ids []uint) ([]*Tag, error) {
	var tags []*Tag
	err := sql.DB.Where("id IN (?)", ids).Find(&tags).Error

	return tags, err
}


func (ts *tagService) FindByTopicID(id uint) ([]*Tag, error) {
	var tags []*Tag
	query := `select t.id, t.name from tags t
					inner join topic_tags tt 
					on tt.tag_id = t.id and tt.topic_id = ?`
	err := sql.DB.Raw(query, id).Scan(&tags).Error
	if err != nil {
		return nil, err
	}
	
	return tags, nil
}

func (ts *tagService) FindByName(name string) ([]*Tag, error) {
	var tags []*Tag

	err := sql.DB.Where("name LIKE ?", "%" + name + "%").Find(&tags).Error

	return tags, err
}

func (ts *tagService) FindOneByName(name string) (Tag, error) {
	var tag Tag

	err := sql.DB.Where("name = ?", name).Find(&tag).Error
	return tag, err
}

func (ts *tagService) FindByTopics(topics []*Topic)(error) {
	for _, topic := range topics {
		tags, err := ts.FindByTopicID(topic.ID)
		if err != nil {
			return err
		}

		topic.Tags = tags
	}

	return nil
}

func (ts *tagService) Create(tag *Tag) error {
	return sql.DB.Create(tag).Error
}

func (ts *tagService) DeleteByTopicID(id uint) error {
	return sql.DB.Delete(&TopicTag{}, "topic_id = ?", id).Error
}
