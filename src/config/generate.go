package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// CfgFileName is the file name of config file
const (
	QueenCfgFileName  = "beehive-queen.conf.yaml"
	WorkerCfgFileName = "beehive-worker.conf.yaml"
)

// roles of running as
const (
	QueenRole  = "queen"
	WorkerRole = "worker"
)

// errors
var (
	ErrUnsupportedRole = fmt.Errorf("unsupported role")
)

// Init generates a config file
func Init(role string) error {
	switch role {
	case QueenRole:
		return initQueenCfgFile()
	case WorkerRole:
		return initWorkerCfgFile()
	default:
		return ErrUnsupportedRole
	}
}

func initQueenCfgFile() error {
	path := fmt.Sprintf("./%s", QueenCfgFileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tpl, err := newDefaultQueenCfg()
	if err != nil {
		return err
	}

	_, err = file.Write(tpl)
	if err != nil {
		return err
	}
	file.Sync()

	return nil
}

func newDefaultQueenCfg() ([]byte, error) {
	db := &DB{
		Type:    "mongodb",
		MongoDB: "mongodb://<username>:<password>@127.0.0.1:27017/<dbname>",
	}

	cfg := &Queen{
		DB: db,
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func initWorkerCfgFile() error {
	path := fmt.Sprintf("./%s", WorkerCfgFileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tpl, err := newDefaultWorkerCfg()
	if err != nil {
		return err
	}

	_, err = file.Write(tpl)
	if err != nil {
		return err
	}
	file.Sync()

	return nil
}

func newDefaultWorkerCfg() ([]byte, error) {
	queen := &QueenServer{
		Host: "http://127.0.0.1:13000",
	}

	rule := &Rule{
		Tag:       "tag",
		RegexpTpl: "^.*$",
		Desc:      "example rule",
	}

	rules := make([]*Rule, 0)
	rules = append(rules, rule)

	cfg := &Worker{
		LogSource: "path/to/log/file",
		Rules:     rules,
		Queen:     queen,
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	return data, nil
}
