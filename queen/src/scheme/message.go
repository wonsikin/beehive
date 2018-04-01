package scheme

import "time"

// Message represents a message which will be stored in the DB
type Message struct {
	Tag        string     `json:"tag" bson:"tag"`
	Text       string     `json:"text" bson:"text"`
	Desc       string     `json:"desc" bson:"desc"`
	CreateTime *time.Time `json:"createTime" bson:"createTime"`
}
