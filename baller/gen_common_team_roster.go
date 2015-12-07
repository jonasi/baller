package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_common_team_roster(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("common_team_roster", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.CommonTeamRosterOptions
	)

	fs.IntVar(&options.TeamID, "TeamID", 0, "")
	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.Season, "Season", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.CommonTeamRoster(&options)
}

func init() {
	methods["common_team_roster"] = cmd_common_team_roster
}
