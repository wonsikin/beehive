package config

import (
	"io/ioutil"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// Worker represents a configuration of worker
type Worker struct {
	LogSource string       `yaml:"logSource"` // the log file the worker digging
	Rules     []*Rule      `yaml:"rules"`     // xxxx
	Queen     *QueenServer `yaml:"queen"`
}

// QueenServer represents beehive-queen's information
type QueenServer struct {
	Host string `yaml:"host"`
}

// Rule represents a log watch rule
type Rule struct {
	Tag       string         `yaml:"tag"`    // rule tag
	RegexpTpl string         `yaml:"regexp"` // regexp to seek log
	Regexp    *regexp.Regexp `yaml:"-"`      // regexp
	Desc      string         `yaml:"desc"`   // description of rule
}

// ParseWorkerCfg parses the config file. It returns an instance of Worker configuration.
func ParseWorkerCfg(path string) (*Worker, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse config file and return
	cfg := &Worker{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	for _, rule := range cfg.Rules {
		re, err := regexp.Compile(rule.RegexpTpl)
		if err != nil {
			return nil, err
		}

		rule.Regexp = re
	}

	return cfg, nil
}
