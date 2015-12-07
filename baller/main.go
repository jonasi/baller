package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/jonasi/baller"
)

var methods = map[string]func(*baller.Client) (interface{}, error){}

func main() {
	cl := baller.New()

	if len(os.Args) == 1 {
		usage()
		os.Exit(0)
	}

	fn, ok := methods[os.Args[1]]

	if !ok {
		fmt.Fprintf(os.Stderr, "Invalid method %s\n\n", os.Args[1])
		usage()
		os.Exit(1)
	}

	data, err := fn(cl)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(data, "", "   ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func usage() {
	fmt.Fprintln(os.Stderr, `Available Methods:`)

	cmds := make([]string, len(methods))

	i := 0
	for k := range methods {
		cmds[i] = k
		i++
	}

	sort.Strings(cmds)

	for _, k := range cmds {
		fmt.Fprintf(os.Stderr, "\t%s\n", k)
	}
}
