package models

import (
	sql "groot/db"
)

type TopicTag struct {
	BaseModel
	TopicID			uint			`json:"topicID"`
	TagID				uint			`json:"tagID"`
}

func (tt *TopicTag) GroupBy(uid uint) (tags []*Tag, err error) {
	fields := "tag.*, count(*) as count"
	joinsTopic := "JOIN topics t ON t.id = tt.topic_id AND t.author_id = ?"
	joinsTag	 := "JOIN tags tag ON tt.tag_id = tag.id"

	err = sql.DB.Table("topic_tags tt").Select(fields).Joins(joinsTopic, uid).Joins(joinsTag).Group("tt.tag_id").Order("t.created_at ASC").Scan(&tags).Error
	return
}
