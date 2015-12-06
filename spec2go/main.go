package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jonasi/baller/spec"
	"github.com/serenize/snaker"
	"os"
	"strings"
	"text/template"
)

var specFile = flag.String("spec", "", "")

func main() {
	flag.Parse()

	if err := do(*specFile); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func do(path string) error {
	f := os.Stdin

	if path != "" && path != "-" {
		var err error
		if f, err = os.Open(path); err != nil {
			return err
		}
	}

	var ep spec.Endpoint

	if err := json.NewDecoder(f).Decode(&ep); err != nil {
		return err
	}

	return apiTemplate.Execute(os.Stdout, map[string]interface{}{
		"spec": ep,
	})
}

var funcs = template.FuncMap{
	"responseName": func(str string) string {
		return strings.Title(str) + "Response"
	},
	"optionsName": func(str string) string {
		return strings.Title(str) + "Options"
	},
	"methodName": strings.Title,
	"rsName": func(sp spec.Endpoint, rs spec.ResultSet) string {
		return strings.Title(sp.Name) + rs.Name
	},
	"encodeFunc": func(str string) string {
		return "encode" + strings.Title(str)
	},
	"fieldName": func(str string) string {
		return snaker.SnakeToCamel(strings.ToLower(str))
	},
}

var apiTemplate = template.Must(template.New("").Funcs(funcs).Parse(`
{{ $spec := .spec }}
{{ range $i, $typ := .spec.ResultSets }}
type {{ rsName $spec $typ }} struct {
	{{ range $k, $p := $typ.Values }}
	{{ fieldName $p.Name }} {{ $p.Type }} ` + "`header:" + `"{{ $p.Name }}"` + "`" + `{{ end }}
}
{{ end }}

type {{ responseName $spec.Name }} struct {
	{{ range $j, $rs := $spec.ResultSets }}
	{{ $rs.Name }} []{{ rsName $spec $rs }}{{ end }}
}

type {{ optionsName $spec.Name }} struct {
	{{ range $j, $rs := $spec.Parameters }}
	{{ $rs.Name }} {{ $rs.Type }}{{ end }}
}

func (c *Client) {{ methodName $spec.Name }}(options *{{ optionsName $spec.Name }}) (*{{ responseName $spec.Name }}, error) {
	var (
		q = url.Values{}
		url = baseURL + "{{ $spec.Name }}?"
		dest  {{ responseName $spec.Name }}
		res result
	)

	{{ range $j, $param := $spec.Parameters }}
		q.Set("{{ $param.Name }}", {{ encodeFunc $param.Type }}(options.{{ $param.Name }})){{ end }}

	if err := c.do(url + q.Encode(), &res); err != nil {
		return nil, err
	}

	{{ range $j, $rs := $spec.ResultSets }}
	if d, err := res.unmarshalResultSet("{{ $rs.Name }}", {{ rsName $spec $rs }}{}); err == nil {
		dest.{{ $rs.Name }} = d.([]{{ rsName $spec $rs }})
	} else {
		return nil, err
	}
	{{ end }}

	return &dest, nil
}`))
