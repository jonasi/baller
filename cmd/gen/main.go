package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

type param struct {
	Name string
	Type string
}

func (p param) StringEnc() string {
	switch p.Type {
	case "string":
		return p.Name
	case "int":
		return "strconv.Itoa(" + p.Name + ")"
	default:
		panic("Invalid type " + p.Type)
	}
}

func (p param) Flag(setVar string) string {
	switch p.Type {
	case "string":
		return fmt.Sprintf(`%s.String("%s", "", "")`, setVar, p.Name)
	case "int":
		return fmt.Sprintf(`%s.Int("%s", 0, "")`, setVar, p.Name)
	default:
		panic("Invalid type " + p.Type)
	}
}

type method struct {
	Name       string
	Path       string
	Return     string
	Parameters []param
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

	var methods []method

	if err := json.Unmarshal(b, &methods); err != nil {
		panic(err)
	}

	apiFile, err := os.OpenFile(*genFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		panic(err)
	}

	if *mode == "api" {
		err = apiTemplate.Execute(apiFile, map[string]interface{}{
			"methods": methods,
		})
	} else {
		err = cliTemplate.Execute(apiFile, map[string]interface{}{
			"methods": methods,
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
	"strconv"
)

{{ range $i, $method := .methods }}
func (c *Client) {{ $method.Name }}({{ range $j, $param := $method.Parameters }}{{ $param.Name }} {{ $param.Type }}{{ if ne $j ($method.Parameters | len) }}, {{end}}{{ end }}) ({{ $method.Return }}, error) {
	var (
		q = url.Values{}
		url = baseURL + "{{ $method.Path }}?"
		dest  {{ $method.Return }}
	)

	{{ range $j, $param := $method.Parameters }}
		q.Set("{{ $param.Name }}", {{ $param.StringEnc }}){{ end }}

	if err := c.do(url + q.Encode(), &dest); err != nil {
		return nil, err
	}

	return dest, nil
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
	{{ range $i, $method := .methods }}"{{ $method.Name }}": {
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
