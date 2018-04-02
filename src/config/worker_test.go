package config

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseWorkerCfg(t *testing.T) {
	Convey("TestParseWorkerCfg", t, func() {
		err := initWorkerCfgFile()
		So(err, ShouldBeNil)

		path := fmt.Sprintf("./%s", WorkerCfgFileName)
		cfg, err := ParseWorkerCfg(path)
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
	})
}
