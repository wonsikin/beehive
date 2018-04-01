package db

import (
	"errors"

	"github.com/CardInfoLink/log"

	"github.com/wonsikin/beehive/queen/src/cfg"
	"github.com/wonsikin/beehive/queen/src/db/mongo"
)

// database type constants
const (
	dbMongo = "mongodb"
)

// error variables
var (
	ErrUnsupportedDBType = errors.New("unsupported db type")
)

// Connect connects to database
func Connect(config *cfg.DB) error {
	log.Debugf("cfg is %#+v", config)
	switch config.Type {
	case dbMongo:
		err := mongo.Init(config.MongoDB)
		if err != nil {
			return err
		}
	default:
		return ErrUnsupportedDBType
	}
	return nil
}
