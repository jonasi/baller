package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_play_by_play_v2(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("play_by_play_v2", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.PlayByPlayV2Options
	)

	fs.IntVar(&options.EndPeriod, "EndPeriod", 0, "")
	fs.StringVar(&options.GameID, "GameID", "", "")
	fs.IntVar(&options.StartPeriod, "StartPeriod", 0, "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.PlayByPlayV2(&options)
}

func init() {
	methods["play_by_play_v2"] = cmd_play_by_play_v2
}
