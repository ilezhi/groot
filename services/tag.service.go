package services

import (
	. "groot/db"
	. "groot/models"
)

type ITag interface {
	Find() ([]*Tag, error)
	FindByID(id uint) (*Tag, error)
	FindByIDs(ids []uint) ([]*Tag, error)
	FindByTopicID(id uint) (*Tag, error)
	FindByName(name string) ([]*Tag, error)
	Create(tag *Tag) error
}

type tagService struct{}

var TagService = tagService{}

func (ts *tagService) Find() ([]*Tag, error) {
	var tags []*Tag
	err := DB.Find(&tags).Error

	return tags, err
}

func (ts *tagService) FindByID(id uint) (*Tag, error) {
	var tag Tag

	err := DB.Find(&tag, id).Error

	return &tag, err
}

func (ts *tagService) FindByIDs(ids []uint) ([]*Tag, error) {
	var tags []*Tag
	err := DB.Where("id IN (?)", ids).Find(&tags).Error

	return tags, err
}

func (ts *tagService) FindByTopicID(id uint) ([]*Tag, error) {
	var tags []*Tag
	sql := `select t.id, t.name from tags t
					inner join topic_tags tt 
					on tt.tag_id = t.id and tt.topic_id = ?`
	err := DB.Raw(sql, id).Scan(&tags).Error
	if err != nil {
		return nil, err
	}
	
	return tags, nil
}

func (ts *tagService) FindByName(name string) ([]*Tag, error) {
	var tags []*Tag

	err := DB.Where("name LIKE ?", name).Find(&tags).Error

	return tags, err
}

func (ts *tagService) Create(tag *Tag) error {
	err := DB.Create(tag).Error

	if err != nil {
		return err
	}

	return nil
}
