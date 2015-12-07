package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_franchise_history(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("franchise_history", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.FranchiseHistoryOptions
	)

	fs.StringVar(&options.LeagueID, "LeagueID", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.FranchiseHistory(&options)
}

func init() {
	methods["franchise_history"] = cmd_franchise_history
}
