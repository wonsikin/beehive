package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCfgFileInit(t *testing.T) {
	Convey("TestCfgFileInit", t, func() {
		wFile := fmt.Sprintf("./%s", WorkerCfgFileName)
		qFile := fmt.Sprintf("./%s", QueenCfgFileName)
		os.Remove(wFile)
		os.Remove(qFile)

		role := ""
		err := Init(role)
		So(err, ShouldEqual, ErrUnsupportedRole)

		role = QueenRole
		err = Init(role)
		So(err, ShouldBeNil)
		_, err = ioutil.ReadFile(qFile)
		So(err, ShouldBeNil)

		role = WorkerRole
		err = Init(role)
		So(err, ShouldBeNil)
		_, err = ioutil.ReadFile(wFile)
		So(err, ShouldBeNil)
	})
}

func TestNewDefaultWorkerCfg(t *testing.T) {
	Convey("TestNewDefaultWorkerCfg", t, func() {
		data, err := newDefaultWorkerCfg()
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		So(len(data), ShouldBeGreaterThan, 0)
	})
}

func TestNewDefaultQueenCfg(t *testing.T) {
	Convey("TestNewDefaultQueenCfg", t, func() {
		data, err := newDefaultQueenCfg()
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		So(len(data), ShouldBeGreaterThan, 0)
	})
}

func TestInitWorkerCfgFile(t *testing.T) {
	Convey("TestInitWorkerCfgFile", t, func() {
		path := fmt.Sprintf("./%s", WorkerCfgFileName)
		os.Remove(path)

		data, err := newDefaultWorkerCfg()
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		So(len(data), ShouldBeGreaterThan, 0)

		err = initWorkerCfgFile()
		So(err, ShouldBeNil)

		ndata, err := ioutil.ReadFile(path)
		So(err, ShouldBeNil)
		So(len(data), ShouldEqual, len(ndata))
		So(string(data), ShouldEqual, string(ndata))
	})
}

func TestInitQueenCfgFile(t *testing.T) {
	Convey("TestInitQueenCfgFile", t, func() {
		path := fmt.Sprintf("./%s", QueenCfgFileName)
		os.Remove(path)

		data, err := newDefaultQueenCfg()
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		So(len(data), ShouldBeGreaterThan, 0)

		err = initQueenCfgFile()
		So(err, ShouldBeNil)

		ndata, err := ioutil.ReadFile(path)
		So(err, ShouldBeNil)
		So(len(data), ShouldEqual, len(ndata))
		So(string(data), ShouldEqual, string(ndata))
	})
}
