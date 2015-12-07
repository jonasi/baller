package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_team_game_log(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("team_game_log", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.TeamGameLogOptions
	)

	fs.IntVar(&options.TeamID, "TeamID", 0, "")
	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.Season, "Season", "", "")
	fs.StringVar(&options.SeasonType, "SeasonType", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.TeamGameLog(&options)
}

func init() {
	methods["team_game_log"] = cmd_team_game_log
}
