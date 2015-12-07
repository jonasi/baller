package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_team_info_common(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("team_info_common", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.TeamInfoCommonOptions
	)

	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.Season, "Season", "", "")
	fs.StringVar(&options.SeasonType, "SeasonType", "", "")
	fs.IntVar(&options.TeamID, "TeamID", 0, "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.TeamInfoCommon(&options)
}

func init() {
	methods["team_info_common"] = cmd_team_info_common
}
