package cfg

import (
	"fmt"
	"os"
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

	_, err = file.WriteString(tpl)
	if err != nil {
		return err
	}
	file.Sync()

	return nil
}

const tpl = `# beehive-queen configuration
db:
	type: 
	mongo:
`
