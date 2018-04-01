package cfg

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// CfgFileName is the file name of config file
const CfgFileName = "beehive-queen.conf.yaml"

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
	db := &DB{
		Type:    "",
		MongoDB: "",
	}

	cfg := &Config{
		DB: db,
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	return data, nil
}
