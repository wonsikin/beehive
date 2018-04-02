package db

import (
	"errors"

	"github.com/wonsikin/beehive/src/config"
	"github.com/wonsikin/beehive/src/db/mongo"
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
func Connect(cfg *config.DB) error {
	switch cfg.Type {
	case dbMongo:
		err := mongo.Init(cfg.MongoDB)
		if err != nil {
			return err
		}
	default:
		return ErrUnsupportedDBType
	}
	return nil
}
