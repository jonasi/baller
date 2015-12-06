package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jonasi/baller/spec"
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

	return cliTemplate.Execute(os.Stdout, map[string]interface{}{
		"spec": ep,
	})
}

var funcs = template.FuncMap{
	"methodName": strings.Title,
	"optionsName": func(str string) string {
		return strings.Title(str) + "Options"
	},
	"flag": func(typ, name, setVar, varName string) string {
		switch typ {
		case "bool":
			return fmt.Sprintf(`%s.BoolVar(&%s, "%s", false, "")`, setVar, varName, name)
		case "string":
			return fmt.Sprintf(`%s.StringVar(&%s, "%s", "", "")`, setVar, varName, name)
		case "int":
			return fmt.Sprintf(`%s.IntVar(&%s, "%s", 0, "")`, setVar, varName, name)
		default:
			panic("Invalid type " + typ)
		}
	},
}

var cliTemplate = template.Must(template.New("").Funcs(funcs).Parse(`
{{ $spec := .spec }}
func cmd_{{ $spec.Name }}(cl *baller.Client) (interface{}, error) {
	var (
		fs = flag.NewFlagSet("{{ $spec.Name }}", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.{{ optionsName $spec.Name }}
	)

	{{ range $j, $param := $spec.Parameters }}{{ (flag $param.Type $param.Name "fs" (print "options." $param.Name) ) }}
	{{ end }}

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.{{ methodName $spec.Name }}(&options)
}

func init() {
	methods["{{ $spec.Name }}"] = cmd_{{ $spec.Name }}
}
`))
