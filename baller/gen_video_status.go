package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_video_status(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("video_status", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.VideoStatusOptions
	)

	fs.StringVar(&options.LeagueID, "LeagueID", "", "")
	fs.StringVar(&options.GameDate, "GameDate", "", "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.VideoStatus(&options)
}

func init() {
	methods["video_status"] = cmd_video_status
}
