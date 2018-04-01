package cfg

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config represents a configuration of worker
type Config struct {
	DB *DB `yaml:"db"` // database configuration
}

// DB represents a database configuration
type DB struct {
	Type    string `yaml:"type"`    // the database the worker store the messages
	MongoDB string `yaml:"mongodb"` // mongodb connenct URL. format [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
}

// Parse parses the config file and returns a instance of Config
func Parse(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse config file and return
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
