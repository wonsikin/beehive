package mongo

import (
	"time"

	"github.com/CardInfoLink/log"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

// Init inits the mongoDB connection
func Init(url string) error {
	session, err := mgo.DialWithTimeout(url, 15*time.Second)
	if err != nil {
		return err
	}

	db = session.DB("settle")
	session.SetMode(mgo.Eventual, true) // 最终一致性即可，读写分离
	session.SetSafe(&mgo.Safe{})
	log.Info("connect to mongodb successfully")
	return nil
}
