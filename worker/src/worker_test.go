package src

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wonsikin/beehive/worker/src/cfg"
)

func TestNewWorker(t *testing.T) {
	Convey("Test new worker", t, func() {
		host := "http://127.0.0.1:13000"
		q := &cfg.Queen{
			Host: host,
		}
		config := &cfg.Config{
			Queen: q,
		}

		worker, err := NewWorker(config)
		So(err, ShouldBeNil)
		So(worker, ShouldNotBeNil)
		So(worker.Config, ShouldNotBeNil)
		So(worker.System, ShouldNotBeNil)
		So(worker.HTTPClient, ShouldNotBeNil)
		So(worker.HTTPClient.Host, ShouldEqual, host)
	})
}

func TestNewMsgPayload(t *testing.T) {
	Convey("Test new worker", t, func() {
		host := "http://127.0.0.1:13000"
		q := &cfg.Queen{
			Host: host,
		}
		config := &cfg.Config{
			Queen: q,
		}

		worker, err := NewWorker(config)
		So(err, ShouldBeNil)
		So(worker, ShouldNotBeNil)
		msg := worker.newMsgPayload("log content")
		So(msg, ShouldNotBeNil)
		So(msg.LogContent, ShouldEqual, "log content")
	})
}
