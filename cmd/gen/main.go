package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/serenize/snaker"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type spec struct {
	Methods     []method     `json:"methods"`
	ResultTypes []resultType `json:"result_types"`
}

type method struct {
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Parameters []param     `json:"parameters"`
	Results    []resultRef `json:"results"`
}

type resultRef struct {
	Name       string `json:"name"`
	ResultType string `json:"result_type"`
}

type resultType struct {
	Name    string  `json:"name"`
	Headers []param `json:"headers"`
}

type param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (p param) FieldName() string {
	return snaker.SnakeToCamel(strings.ToLower(p.Name))
}

func (p param) StringEnc() string {
	switch p.Type {
	case "string", "int", "bool":
		return fmt.Sprintf("encode%s(%s)", strings.Title(p.Type), p.Name)
	default:
		panic("Invalid type " + p.Type)
	}
}

func (p param) Flag(setVar string) string {
	switch p.Type {
	case "bool":
		return fmt.Sprintf(`%s.Bool("%s", false, "")`, setVar, p.Name)
	case "string":
		return fmt.Sprintf(`%s.String("%s", "", "")`, setVar, p.Name)
	case "int":
		return fmt.Sprintf(`%s.Int("%s", 0, "")`, setVar, p.Name)
	default:
		panic("Invalid type " + p.Type)
	}
}

var (
	specFile = flag.String("spec", "", "")
	mode     = flag.String("mode", "", "")
	genFile  = flag.String("genfile", "", "")
)

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*specFile)

	if err != nil {
		panic(err)
	}

	var spec spec

	if err := json.Unmarshal(b, &spec); err != nil {
		panic(err)
	}

	apiFile, err := os.OpenFile(*genFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		panic(err)
	}

	if *mode == "api" {
		err = apiTemplate.Execute(apiFile, map[string]interface{}{
			"spec": spec,
		})
	} else {
		err = cliTemplate.Execute(apiFile, map[string]interface{}{
			"spec": spec,
		})
	}

	if err != nil {
		panic(err)
	}

	apiFile.Close()

	if err := exec.Command("go", "fmt", *genFile).Run(); err != nil {
		panic(err)
	}
}

var apiTemplate = template.Must(template.New("").Parse(
	`//go:generate go run cmd/gen/main.go -spec api_spec.json -genfile=$GOFILE -mode=api

package baller

import (
	"net/url"
)

{{ range $i, $typ := .spec.ResultTypes }}
type Result{{ $typ.Name }} struct {
	{{ range $k, $p := $typ.Headers }}
	{{ $p.FieldName }} {{ $p.Type }} ` + "`header:" + `"{{ $p.Name }}"` + "`" + `{{ end }}
}
{{ end }}

{{ range $i, $method := .spec.Methods }}
type {{ $method.Name }}Response struct {
	{{ range $j, $res := $method.Results }}
	{{ $res.Name }} []Result{{ $res.ResultType }}{{ end }}
}

func (c *Client) {{ $method.Name }}({{ range $j, $param := $method.Parameters }}{{ $param.Name }} {{ $param.Type }}{{ if ne $j ($method.Parameters | len) }}, {{end}}{{ end }}) (*{{ $method.Name }}Response, error) {
	var (
		q = url.Values{}
		url = baseURL + "{{ $method.Path }}?"
		dest  {{ $method.Name }}Response
		res result
	)

	{{ range $j, $param := $method.Parameters }}
		q.Set("{{ $param.Name }}", {{ $param.StringEnc }}){{ end }}

	if err := c.do(url + q.Encode(), &res); err != nil {
		return nil, err
	}

	{{ range $j, $res := $method.Results }}
	if d, err := res.unmarshalResultSet("{{ $res.Name }}", Result{{ $res.ResultType }}{}); err == nil {
		dest.{{ $res.Name }} = d.([]Result{{ $res.ResultType }})
	} else {
		return nil, err
	}
	{{ end }}

	return &dest, nil
}
{{ end }}`))

var cliTemplate = template.Must(template.New("").Parse(
	`//go:generate go run ../gen/main.go -spec ../../api_spec.json -genfile=$GOFILE -mode=cli

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jonasi/baller"
)

var methods = map[string]struct{
	Do func(*baller.Client) (interface{}, error)
}{
	{{ range $i, $method := .spec.Methods }}"{{ $method.Name }}": {
		Do: func(cl *baller.Client) (interface{}, error) {
			var (
				fs = flag.NewFlagSet("{{ $method.Name }}", flag.ExitOnError)
				verbose = fs.Bool("verbose", false, "")
				{{ range $j, $param := $method.Parameters }}{{ $param.Name }} = {{ ($param.Flag "fs") }}
				{{ end }}
			)

			fs.Parse(os.Args[2:])

			if *verbose {
				cl.Logger = os.Stderr
			}

			return cl.{{ $method.Name }}({{ range $j, $param := $method.Parameters }}*{{ $param.Name }}{{ if ne $j ($method.Parameters | len) }}, {{ end }}{{ end }})
		},
},
	{{ end }}
}

func main() {
	cl := baller.New()
	data, err := methods[os.Args[1]].Do(cl)

	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(data, "", "   ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
`))
