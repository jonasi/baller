package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_boxscore_traditional_v2(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("boxscore_traditional_v2", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.BoxscoreTraditionalV2Options
	)

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.BoxscoreTraditionalV2(&options)
}

func init() {
	methods["boxscore_traditional_v2"] = cmd_boxscore_traditional_v2
}
