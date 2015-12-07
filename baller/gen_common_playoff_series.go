package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_common_playoff_series(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("common_playoff_series", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.CommonPlayoffSeriesOptions
	)

	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.Season, "Season", "", "")
	fs.StringVar(&options.SeriesID, "SeriesID", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.CommonPlayoffSeries(&options)
}

func init() {
	methods["common_playoff_series"] = cmd_common_playoff_series
}
