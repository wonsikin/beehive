package mongo

import (
	"time"

	"github.com/wonsikin/beehive/queen/src/scheme"
)

var msgCollectionName = "hive.msg"

// MsgCollection represents a message collection in the mongo db
type MsgCollection struct {
	name string
}

// NewMsgCollection returns a new message collection
func NewMsgCollection() *MsgCollection {
	return &MsgCollection{
		name: msgCollectionName,
	}
}

// Insert inserts one record into the msg
func (c *MsgCollection) Insert(msg *scheme.Message) error {
	now := time.Now()
	msg.CreateTime = &now
	return db.C(c.name).Insert(msg)
}
