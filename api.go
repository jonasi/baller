//go:generate go run cmd/gen/main.go -spec api_spec.json -genfile=$GOFILE -mode=api

package baller

import (
	"net/url"
	"strconv"
)

func (c *Client) Scoreboard(GameDate string, LeagueID string, DayOffset int) (*Scoreboard, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboard?"
		dest *Scoreboard
	)

	q.Set("GameDate", GameDate)
	q.Set("LeagueID", LeagueID)
	q.Set("DayOffset", strconv.Itoa(DayOffset))

	if err := c.do(url+q.Encode(), &dest); err != nil {
		return nil, err
	}

	return dest, nil
}

func (c *Client) ScoreboardV2(GameDate string, LeagueID string, DayOffset int) (*ScoreboardV2, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboardv2?"
		dest *ScoreboardV2
	)

	q.Set("GameDate", GameDate)
	q.Set("LeagueID", LeagueID)
	q.Set("DayOffset", strconv.Itoa(DayOffset))

	if err := c.do(url+q.Encode(), &dest); err != nil {
		return nil, err
	}

	return dest, nil
}
