package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_scoreboard_v2(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("scoreboard_v2", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.ScoreboardV2Options
	)

	fs.StringVar(&options.GameDate, "GameDate", "", "")
	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.DayOffset, "DayOffset", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.ScoreboardV2(&options)
}

func init() {
	methods["scoreboard_v2"] = cmd_scoreboard_v2
}
