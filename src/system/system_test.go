package system

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCurrentDir(t *testing.T) {
	Convey("Test current directory", t, func() {
		dir, err := currentDir()
		So(err, ShouldBeNil)
		So(dir, ShouldNotBeNil)
		So(len(dir), ShouldBeGreaterThan, 0)
	})
}

func TestLogPath(t *testing.T) {
	Convey("Test log path", t, func() {
		dir, err := currentDir()
		So(err, ShouldBeNil)
		So(dir, ShouldNotBeNil)
		So(len(dir), ShouldBeGreaterThan, 0)

		path := "/path/to/log"
		lp, err := logPath(path)
		So(err, ShouldBeNil)
		So(lp, ShouldNotBeNil)
		So(lp, ShouldEqual, path)

		path = "./path/to/log"
		lp, err = logPath(path)
		So(err, ShouldBeNil)
		So(lp, ShouldNotBeNil)
		So(lp, ShouldNotEqual, path)
		So(lp, ShouldEqual, fmt.Sprintf("%s%s", dir, path[1:]))
	})
}

func TestNewInformation(t *testing.T) {
	Convey("Test new system information", t, func() {
		path := "./path/to/log"
		info, err := NewInformation(path)
		So(err, ShouldBeNil)
		So(info, ShouldNotBeNil)
		So(info.HostName, ShouldNotBeEmpty)
		So(info.IP, ShouldNotBeEmpty)
		So(info.LogPath, ShouldNotBeEmpty)
	})
}
