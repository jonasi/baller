package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_boxscore_advanced_v2(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("boxscore_advanced_v2", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.BoxscoreAdvancedV2Options
	)

	fs.StringVar(&options.GameID, "GameID", "", "")
	fs.IntVar(&options.StartPeriod, "StartPeriod", 0, "")
	fs.IntVar(&options.EndPeriod, "EndPeriod", 0, "")
	fs.IntVar(&options.StartRange, "StartRange", 0, "")
	fs.IntVar(&options.EndRange, "EndRange", 0, "")
	fs.IntVar(&options.RangeType, "RangeType", 0, "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.BoxscoreAdvancedV2(&options)
}

func init() {
	methods["boxscore_advanced_v2"] = cmd_boxscore_advanced_v2
}
