package model

type TopicTag struct {
	BaseModel
	TopicID			uint			`json:"topicID"`
	TagID				uint			`json:"tagID"`
}
