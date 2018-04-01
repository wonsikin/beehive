package cfg

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	yaml "gopkg.in/yaml.v2"
)

func TestNewDefaultConfigTemplate(t *testing.T) {
	Convey("Test new default config template", t, func() {
		tpl, err := newDefaultConfigTemplate()
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)
		So(len(tpl), ShouldBeGreaterThan, 0)
	})
}

func TestWorkerInit(t *testing.T) {
	Convey("", t, func() {
		err := Init()
		So(err, ShouldBeNil)
		tpl, err := newDefaultConfigTemplate()
		So(err, ShouldBeNil)
		So(tpl, ShouldNotBeNil)

		data, err := ioutil.ReadFile(fmt.Sprintf("./%s", CfgFileName))
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		So(len(data), ShouldEqual, len(tpl))
		So(string(data), ShouldEqual, string(tpl))

		config := &Config{}
		err = yaml.Unmarshal(data, config)
		So(err, ShouldBeNil)
	})
}
