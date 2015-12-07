package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_scoreboard(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("scoreboard", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.ScoreboardOptions
	)

	fs.StringVar(&options.GameDate, "GameDate", "", "")
	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.DayOffset, "DayOffset", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.Scoreboard(&options)
}

func init() {
	methods["scoreboard"] = cmd_scoreboard
}
