package baller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
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

type result struct {
	Resource   string
	Parameters map[string]interface{}
	ResultSets []struct {
		Name    string
		Headers []string
		RowSet  []json.RawMessage
	}
}

func (r *result) unmarshalResultSet(name string, typ interface{}) (interface{}, error) {
	var (
		headers []string
		row     []json.RawMessage
	)

	for i := range r.ResultSets {
		if r.ResultSets[i].Name == name {
			headers = r.ResultSets[i].Headers
			row = r.ResultSets[i].RowSet
			break
		}
	}

	if headers == nil {
		return nil, fmt.Errorf("Row set %s not found", name)
	}

	dest := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(typ)), len(row), len(row)).Interface()

	if err := decodeResultSet(dest, headers, row); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *result) unmarshalResultSets(m map[string]interface{}) (map[string]interface{}, error) {
	ret := map[string]interface{}{}

	for k, typ := range m {
		v, err := r.unmarshalResultSet(k, typ)

		if err != nil {
			return nil, err
		}

		ret[k] = v
	}

	return ret, nil
}
