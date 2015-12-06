package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_boxscore(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("boxscore", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.BoxscoreOptions
	)

	fs.IntVar(&options.StartRange, "StartRange", 0, "")
	fs.IntVar(&options.EndRange, "EndRange", 0, "")
	fs.IntVar(&options.RangeType, "RangeType", 0, "")
	fs.StringVar(&options.GameID, "GameID", "", "")
	fs.IntVar(&options.StartPeriod, "StartPeriod", 0, "")
	fs.IntVar(&options.EndPeriod, "EndPeriod", 0, "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.Boxscore(&options)
}

func init() {
	methods["boxscore"] = cmd_boxscore
}
