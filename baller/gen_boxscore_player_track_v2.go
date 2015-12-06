package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_boxscore_player_track_v2(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("boxscore_player_track_v2", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.BoxscorePlayerTrackV2Options
	)

	fs.StringVar(&options.GameID, "GameID", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.BoxscorePlayerTrackV2(&options)
}

func init() {
	methods["boxscore_player_track_v2"] = cmd_boxscore_player_track_v2
}
