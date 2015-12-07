package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)

func cmd_play_by_play(cl *baller.Client) (interface{}, error) {
	var (
		fs      = flag.NewFlagSet("play_by_play", flag.ExitOnError)
		verbose = fs.Bool("verbose", false, "")
		options baller.PlayByPlayOptions
	)

	fs.StringVar(&options.GameID, "GameID", "", "")
	fs.IntVar(&options.StartPeriod, "StartPeriod", 0, "")
	fs.IntVar(&options.EndPeriod, "EndPeriod", 0, "")

	fs.Parse(os.Args[2:])

	if *verbose {
		cl.Logger = os.Stderr
	}

	return cl.PlayByPlay(&options)
}

func init() {
	methods["play_by_play"] = cmd_play_by_play
}
