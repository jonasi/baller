package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_common_all_players(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("common_all_players", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.CommonAllPlayersOptions
	)

	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.Season, "Season", "", "")
	fs.IntVar(&options.IsOnlyCurrentSeason, "IsOnlyCurrentSeason", 0, "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.CommonAllPlayers(&options)
}

func init() {
	methods["common_all_players"] = cmd_common_all_players
}
