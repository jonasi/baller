package baller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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

type Parameters map[string]interface{}

func (p *Parameters) UnmarshalJSON(b []byte) error {
	if b[0] == '{' {
		p2 := (*map[string]interface{})(p)
		return json.Unmarshal(b, p2)
	}

	var p2 []map[string]interface{}

	if err := json.Unmarshal(b, &p2); err != nil {
		return err
	}

	*p = map[string]interface{}{}

	for _, m := range p2 {
		for k, v := range m {
			(*p)[k] = v
		}
	}

	return nil
}

type result struct {
	Resource   string
	Parameters Parameters
	ResultSets []struct {
		Name    string
		Headers []string
		RowSet  []json.RawMessage
	}
}

func (r *result) unmarshalResultSets(dests map[string]interface{}) error {
	for name, dest := range dests {
		if err := r.unmarshalResultSet(name, dest); err != nil {
			return err
		}
	}

	return nil
}

func (r *result) unmarshalResultSet(name string, dest interface{}) error {
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
		return fmt.Errorf("Row set %s not found", name)
	}

	rv := reflect.ValueOf(dest)
	sl := reflect.MakeSlice(reflect.SliceOf(rv.Type().Elem().Elem()), len(row), len(row))

	rv.Elem().Set(sl)

	if err := decodeResultSet(sl.Interface(), headers, row); err != nil {
		return err
	}

	return nil
}

func decodeResultSet(dest interface{}, headers []string, rowset []json.RawMessage) error {
	var (
		v  = reflect.ValueOf(dest)
		t  = v.Type().Elem()
		mp = mkHeaderMap(t)
		d  = diff(headers, mp)
	)

	if len(d) != 0 {
		dstr := strings.Join(d, ", ")
		return fmt.Errorf("Expected %d headers, found %d for type %T\nDiff: [%s]", len(headers), len(mp), dest, dstr)
	}

	for i := range rowset {
		sl := make([]interface{}, len(headers))

		for j, h := range headers {
			if idx, ok := mp[h]; ok {
				sl[j] = v.Index(i).Field(idx).Addr().Interface()
			}
		}

		if err := json.Unmarshal(rowset[i], &sl); err != nil {
			return err
		}
	}

	return nil
}

func mkHeaderMap(t reflect.Type) map[string]int {
	mp := map[string]int{}

	for i := 0; i < t.NumField(); i++ {
		h := t.Field(i).Tag.Get("header")

		if h == "" {
			continue
		}

		mp[h] = i
	}

	return mp
}

func diff(headers []string, mp map[string]int) []string {
	var (
		exp  = map[string]struct{}{}
		diff = []string{}
	)

	for _, h := range headers {
		if _, ok := mp[h]; !ok {
			diff = append(diff, "Expected "+h)
		}

		exp[h] = struct{}{}
	}

	for h := range mp {
		if _, ok := exp[h]; ok {
			continue
		}

		diff = append(diff, "Unexpected "+h)
	}

	return diff
}
