package models

import (
	sql "groot/db"
)

type TopicTag struct {
	BaseModel
	TopicID			uint			`json:"topicID"`
	TagID				uint			`json:"tagID"`
}


func (tt *TopicTag) GroupByTag(uid uint) (tags []*Tag, err error) {
	fields := "t.*, count(*) as count"
	joinsTopic := "JOIN topics t ON t.id = tt.topic_id"
	joinsUser	 := "JOIN users u ON t.author_id = u.id AND u.id = ?"

	err = sql.DB.Table("topic_tags tt").Select(fields).Joins(joinsTopic).Joins(joinsUser, uid).Group("tt.tag_id").Order("t.created_at ASC").Scan(&tags).Error
	return
}
