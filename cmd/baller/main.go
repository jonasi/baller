//go:generate go run ../gen/main.go -spec ../../api_spec.json -genfile=$GOFILE -mode=cli

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jonasi/baller"
)

var methods = map[string]struct {
	Do func(*baller.Client) (interface{}, error)
}{
	"Scoreboard": {
		Do: func(cl *baller.Client) (interface{}, error) {
			var (
				fs        = flag.NewFlagSet("Scoreboard", flag.ExitOnError)
				verbose   = fs.Bool("verbose", false, "")
				GameDate  = fs.String("GameDate", "", "")
				LeagueID  = fs.String("LeagueID", "", "")
				DayOffset = fs.Int("DayOffset", 0, "")
			)

			fs.Parse(os.Args[2:])

			if *verbose {
				cl.Logger = os.Stderr
			}

			return cl.Scoreboard(*GameDate, *LeagueID, *DayOffset)
		},
	},
	"ScoreboardV2": {
		Do: func(cl *baller.Client) (interface{}, error) {
			var (
				fs        = flag.NewFlagSet("ScoreboardV2", flag.ExitOnError)
				verbose   = fs.Bool("verbose", false, "")
				GameDate  = fs.String("GameDate", "", "")
				LeagueID  = fs.String("LeagueID", "", "")
				DayOffset = fs.Int("DayOffset", 0, "")
			)

			fs.Parse(os.Args[2:])

			if *verbose {
				cl.Logger = os.Stderr
			}

			return cl.ScoreboardV2(*GameDate, *LeagueID, *DayOffset)
		},
	},
}

func main() {
	cl := baller.New()
	data, err := methods[os.Args[1]].Do(cl)

	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(b))
}
