package src

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wonsikin/beehive/src/config"
)

func TestNewWorker(t *testing.T) {
	Convey("Test new worker", t, func() {
		host := "http://127.0.0.1:13000"
		q := &config.QueenServer{
			Host: host,
		}
		cfg := &config.Worker{
			Queen: q,
		}

		worker, err := NewWorker(cfg)
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
		q := &config.QueenServer{
			Host: host,
		}
		cfg := &config.Worker{
			Queen: q,
		}

		worker, err := NewWorker(cfg)
		So(err, ShouldBeNil)
		So(worker, ShouldNotBeNil)
		msg := worker.newMsgPayload("log content")
		So(msg, ShouldNotBeNil)
		So(msg.LogContent, ShouldEqual, "log content")
	})
}
