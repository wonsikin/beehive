package cfg

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewDefaultConfigTemplate(t *testing.T) {
	Convey("Test new default config template", t, func() {
		tpl, err := newDefaultConfigTemplate()
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		So(len(tpl), ShouldBeGreaterThan, 0)
	})
}
