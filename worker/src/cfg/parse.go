package cfg

import (
	"io/ioutil"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// Config represents a configuration of worker
type Config struct {
	LogSource string  `yaml:"logSource"` // the log file the worker digging
	Rules     []*Rule `yaml:"rules"`     // xxxx
	Queen     *Queen  `yaml:"queen"`
}

// Queen represents beehive-queen's information
type Queen struct {
	Host string `yaml:"host"`
}

// Rule represents a log watch rule
type Rule struct {
	Tag       string         `yaml:"tag"`    // rule tag
	RegexpTpl string         `yaml:"regexp"` // regexp to seek log
	Regexp    *regexp.Regexp `yaml:"-"`      // regexp
	Desc      string         `yaml:"desc"`   // description of rule
}

// Parse parses the config file
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

	for _, rule := range cfg.Rules {
		re, err := regexp.Compile(rule.RegexpTpl)
		if err != nil {
			return nil, err
		}

		rule.Regexp = re
	}

	return cfg, nil
}
