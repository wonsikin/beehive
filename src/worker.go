package src

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"

	"github.com/wonsikin/beehive/src/bridge"
	"github.com/wonsikin/beehive/src/config"
	"github.com/wonsikin/beehive/src/scheme"
	"github.com/wonsikin/beehive/src/system"
)

// Worker represents a worker instance
type Worker struct {
	Config     *config.Worker
	System     *system.Information
	HTTPClient *bridge.Client
}

// NewWorker return a new worker
func NewWorker(cfg *config.Worker) (*Worker, error) {
	client := bridge.NewClient(cfg.Queen.Host)

	si, err := system.NewInformation(cfg.LogSource)
	if err != nil {
		return nil, err
	}
	return &Worker{
		Config:     cfg,
		System:     si,
		HTTPClient: client,
	}, nil
}

type msgPayload struct {
	HostName   string `json:"hostName"`
	IP         string `json:"ip"`
	LogPath    string `json:"logPath"`
	LogContent string `json:"logContent"`
}

// Run starts run work
func (w *Worker) newMsgPayload(log string) *msgPayload {
	return &msgPayload{
		HostName:   w.System.HostName,
		IP:         w.System.IP,
		LogPath:    w.System.LogPath,
		LogContent: log,
	}
}

// Run starts run work
func (w *Worker) Run() {
	fmt.Printf("Worker %s is running\n", w.System.HostName)
	seekInfo := tail.SeekInfo{Offset: 0, Whence: os.SEEK_END}
	t, err := tail.TailFile(w.Config.LogSource, tail.Config{Follow: true, Poll: true, ReOpen: true, Location: &seekInfo})
	if err != nil {
		panic(err)
	}

	rules := w.Config.Rules
	for line := range t.Lines {
		for _, rx := range rules {
			if matched := rx.Regexp.MatchString(line.Text); matched {
				pl := w.newMsgPayload(line.Text)

				msg := &scheme.Message{
					Tag:     rx.Tag,
					Payload: pl,
					Desc:    rx.Desc,
				}

				err = w.HTTPClient.SendMessage(msg)
				if err != nil {

				}
				break
			}
		}
	}

	select {}
}
