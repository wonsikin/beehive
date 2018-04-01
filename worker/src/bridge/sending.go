package bridge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/CardInfoLink/log"

	"github.com/wonsikin/beehive/queen/src/scheme"
)

// Client represents a HTTP client
type Client struct {
	Host string
	*http.Client
}

// NewClient returns a new client
func NewClient(host string) *Client {
	return &Client{host, &http.Client{Timeout: 30 * time.Second}}
}

// SendMessage sends message to the queen app
func (c *Client) SendMessage(msg *scheme.Message) error {
	url := fmt.Sprintf("%s%s", c.Host, "/bhq/message")
	body := strings.NewReader(msg.String())
	response, err := c.Post(url, "application/json;charset=UTF-8", body)
	if err != nil {
		log.Errorf("error caught when sending a post request: %s", err)
		return err
	}

	if response.StatusCode == http.StatusOK {
		return nil
	}

	dd, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("error caught when sending a post request: %s", err)
		return err
	}
	defer response.Body.Close()

	errMsg := fmt.Sprintf("request fail: %s-%s", response.Status, string(dd))

	return fmt.Errorf(errMsg)
}
