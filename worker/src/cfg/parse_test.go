package cfg

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	Convey("Test parse config file", t, func() {
		err := Init()
		So(err, ShouldBeNil)

		path := fmt.Sprintf("./%s", CfgFileName)
		cfg, err := Parse(path)
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)

		So(cfg.LogSource, ShouldEqual, "path/to/log/file")
		So(cfg.Queen, ShouldNotBeNil)
		t.Logf("config is %+#v", cfg)
		t.Logf("config is %#+v", cfg.Queen)
	})
}
