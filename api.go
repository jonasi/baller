//go:generate go run cmd/gen/main.go -spec api_spec.json -genfile=$GOFILE -mode=api

package baller

import (
	"net/url"
)

func (c *Client) CommonAllPlayers(LeagueID string, Season string, IsOnlyCurrentSeason bool) (map[string]interface{}, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonallplayers?"
		dest map[string]interface{}
	)

	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("Season", encodeString(Season))
	q.Set("IsOnlyCurrentSeason", encodeBool(IsOnlyCurrentSeason))

	if err := c.do(url+q.Encode(), &dest); err != nil {
		return nil, err
	}

	return dest, nil
}

func (c *Client) Scoreboard(GameDate string, LeagueID string, DayOffset int) (*Scoreboard, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboard?"
		dest *Scoreboard
	)

	q.Set("GameDate", encodeString(GameDate))
	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("DayOffset", encodeInt(DayOffset))

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

	q.Set("GameDate", encodeString(GameDate))
	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("DayOffset", encodeInt(DayOffset))

	if err := c.do(url+q.Encode(), &dest); err != nil {
		return nil, err
	}

	return dest, nil
}
