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
	"unicode"
	"unicode/utf8"
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

func ucfirst(str string) string {
	f, i := utf8.DecodeRuneInString(str)
	return string(unicode.ToUpper(f)) + str[i:]
}

var funcs = template.FuncMap{
	"responseName": func(str string) string {
		return snaker.SnakeToCamel(str) + "Response"
	},
	"optionsName": func(str string) string {
		return snaker.SnakeToCamel(str) + "Options"
	},
	"methodName": snaker.SnakeToCamel,
	"rsName": func(sp, rs string) string {
		return snaker.SnakeToCamel(sp) + ucfirst(rs)
	},
	"encodeFunc": func(str string) string {
		return "encode" + ucfirst(str)
	},
	"fieldName": func(str string) string {
		return ucfirst(snaker.SnakeToCamel(str))
	},
	"fieldNameFromRS": func(str string) string {
		return ucfirst(snaker.SnakeToCamel(strings.ToLower(str)))
	},
}

var apiTemplate = template.Must(template.New("").Funcs(funcs).Parse(`
{{ $spec := .spec }}
type {{ optionsName $spec.Name }} struct {
	{{ range $j, $rs := $spec.Parameters }}
	{{ fieldName $rs.Name }} {{ $rs.Type }}{{ end }}
}

type {{ responseName $spec.Name }} struct {
	{{ range $j, $rs := $spec.ResultSets }}
	{{ fieldName $rs.Name }} []{{ rsName $spec.Name $rs.Name }}{{ end }}
}

func (c *Client) {{ methodName $spec.Name }}(options *{{ optionsName $spec.Name }}) (*{{ responseName $spec.Name }}, error) {
	var (
		q = url.Values{}
		url = baseURL + "{{ $spec.Path }}?"
		dest  {{ responseName $spec.Name }}
		res result
	)

	{{ range $j, $param := $spec.Parameters }}
		q.Set("{{ $param.Name }}", {{ encodeFunc $param.Type }}(options.{{ $param.Name }})){{ end }}

	if err := c.do(url + q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{ {{ range $j, $rs := $spec.ResultSets }}
		"{{ $rs.Name }}": &dest.{{ fieldName $rs.Name }},{{ end }}
	})
	
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

{{ range $i, $rs := .spec.ResultSets }}
type {{ rsName $spec.Name $rs.Name }} struct {
	{{ range $k, $p := $rs.Values }}
	{{ fieldNameFromRS $p.Name }} {{ $p.Type }} ` + "`header:" + `"{{ $p.Name }}"` + "`" + `{{ end }}
}
{{ end }}
`))
