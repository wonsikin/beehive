package config

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseQueenCfg(t *testing.T) {
	Convey("TestParseQueenCfg", t, func() {
		err := initQueenCfgFile()
		So(err, ShouldBeNil)

		path := fmt.Sprintf("./%s", QueenCfgFileName)
		cfg, err := ParseQueenCfg(path)
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
	})
}
