package scheme

import (
	"encoding/json"
	"time"

	"github.com/CardInfoLink/log"
)

// Message represents a message which will be stored in the DB
type Message struct {
	Tag        string      `json:"tag" bson:"tag"`
	Payload    interface{} `json:"payload" bson:"payload"`
	Desc       string      `json:"desc" bson:"desc"`
	CreateTime *time.Time  `json:"createTime" bson:"createTime"`
}

// String returns a JSON string of the struct
func (m *Message) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		log.Errorf("error caught when marshalling data: %s", err)
	}

	return string(data)
}
