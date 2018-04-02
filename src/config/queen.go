package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Queen represents a configuration of queen
type Queen struct {
	DB *DB `yaml:"db"` // database configuration
}

// DB represents a database configuration
type DB struct {
	Type    string `yaml:"type"`    // the database the queen server store the messages
	MongoDB string `yaml:"mongodb"` // mongodb connenct URL. format [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
}

// ParseQueenCfg parses the config file of the queen and returns a instance of Queen configuration
func ParseQueenCfg(path string) (*Queen, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse config file and return
	cfg := &Queen{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
