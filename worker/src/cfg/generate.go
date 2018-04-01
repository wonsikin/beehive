package cfg

import (
	"fmt"
	"os"

	"github.com/CardInfoLink/log"
	yaml "gopkg.in/yaml.v2"
)

// CfgFileName is the file name of config file
const CfgFileName = "beehive-worker.conf.yaml"

// Init generates a config file
func Init() error {
	path := fmt.Sprintf("./%s", CfgFileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tpl, err := newDefaultConfigTemplate()
	if err != nil {
		log.Errorf("error caught when create default config: %s", err)
		return err
	}

	_, err = file.Write(tpl)
	if err != nil {
		return err
	}
	file.Sync()

	return nil
}

func newDefaultConfigTemplate() ([]byte, error) {
	queen := &Queen{
		Host: "http://127.0.0.1:13000",
	}

	rule := &Rule{
		Tag:       "tag",
		RegexpTpl: "^.*$",
		Desc:      "example rule",
	}

	rules := make([]*Rule, 0)
	rules = append(rules, rule)

	cfg := &Config{
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
