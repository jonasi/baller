package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_common_team_years(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("common_team_years", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.CommonTeamYearsOptions
	)

	fs.StringVar(&options.LeagueID, "LeagueID", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.CommonTeamYears(&options)
}

func init() {
	methods["common_team_years"] = cmd_common_team_years
}
