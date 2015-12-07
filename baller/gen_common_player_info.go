package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_common_player_info(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("common_player_info", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.CommonPlayerInfoOptions
	)

	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.IntVar(&options.PlayerID, "PlayerID", 0, "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.CommonPlayerInfo(&options)
}

func init() {
	methods["common_player_info"] = cmd_common_player_info
}
