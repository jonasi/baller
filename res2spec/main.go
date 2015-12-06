package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jonasi/baller/spec"
	"os"
)

var (
	responseFile = flag.String("response", "", "")
	name         = flag.String("name", "", "")
)

func main() {
	flag.Parse()

	err := do(*responseFile, *name)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

type rsVal struct {
	typ string
}

// only care about string,int,float
func (v *rsVal) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		v.typ = "string"
	} else if bytes.Contains(b, []byte(".")) {
		v.typ = "float32"
	} else {
		v.typ = "int"
	}

	return nil
}

type Response struct {
	Resource   string     `json:"resource"`
	Parameters parameters `json:"parameters"`
	ResultSets []struct {
		Name    string    `json:"name"`
		Headers []string  `json:"headers"`
		RowSet  [][]rsVal `json:"rowSet"`
	} `json:"resultSets"`
}

type parameters map[string]rsVal

func (p *parameters) UnmarshalJSON(b []byte) error {
	if b[0] == '{' {
		p2 := (*map[string]rsVal)(p)
		return json.Unmarshal(b, p2)
	}

	var p2 []map[string]rsVal

	if err := json.Unmarshal(b, &p2); err != nil {
		return err
	}

	*p = map[string]rsVal{}

	for _, m := range p2 {
		for k, v := range m {
			(*p)[k] = v
		}
	}

	return nil
}

func do(path, name string) error {
	f := os.Stdin

	if path != "" && path != "-" {
		var err error
		if f, err = os.Open(path); err != nil {
			return err
		}
	}

	var res Response

	if err := json.NewDecoder(f).Decode(&res); err != nil {
		return err
	}

	ep := spec.Endpoint{
		Name:       res.Resource,
		Parameters: make([]spec.Value, len(res.Parameters)),
		ResultSets: make([]spec.ResultSet, len(res.ResultSets)),
	}

	if name != "" {
		ep.Name = name
	}

	i := 0
	for k, v := range res.Parameters {
		ep.Parameters[i].Name = k
		ep.Parameters[i].Type = v.typ
		i++
	}

	for i := range res.ResultSets {
		rs := &res.ResultSets[i]

		ep.ResultSets[i].Name = rs.Name
		ep.ResultSets[i].Values = make([]spec.Value, len(rs.Headers))

		for j, h := range rs.Headers {
			ep.ResultSets[i].Values[j].Name = h
			ep.ResultSets[i].Values[j].Type = rs.RowSet[0][j].typ
		}
	}

	return json.NewEncoder(os.Stdout).Encode(ep)
}
