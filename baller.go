package baller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "http://stats.nba.com/stats/"
)

func New() *Client {
	return &Client{}
}

type Client struct {
	Logger io.Writer
}

func (c *Client) infof(f string, args ...interface{}) {
	if c.Logger != nil {
		fmt.Fprintf(c.Logger, "[INFO] "+f+"\n", args...)
	}
}

func (c *Client) do(url string, dest interface{}) error {
	c.infof("Client request: url=%s", url)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	c.infof("Client Response: code=%d headers=%s", resp.StatusCode, resp.Header)

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		return errors.New(string(b))
	}

	if c.Logger == nil {
		return json.NewDecoder(resp.Body).Decode(dest)
	}

	var b json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&b); err != nil {
		return err
	}

	c.infof("Client Response Body %s", string(b))

	return json.Unmarshal(b, dest)
}
