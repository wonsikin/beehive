package bridge

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wonsikin/beehive/queen/src/scheme"
)

func TestNewClient(t *testing.T) {
	Convey("Test new client", t, func() {
		host := "http://127.0.0.1:13000"
		cl := NewClient(host)
		So(cl, ShouldNotBeNil)
		So(cl.Timeout, ShouldEqual, 30*time.Second)
	})
}

func TestSendingMessage(t *testing.T) {
	Convey("Test sending message", t, func() {
		host := "http://127.0.0.1:13000"
		cl := NewClient(host)
		msg := &scheme.Message{}
		err := cl.SendMessage(msg)
		So(err, ShouldNotBeNil)
	})
}
