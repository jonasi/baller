package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jonasi/baller"
)

var methods = map[string]func(*baller.Client) (interface{}, error){}

func main() {
	cl := baller.New()

	data, err := methods[os.Args[1]](cl)

	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(data, "", "   ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
