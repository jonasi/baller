//go:generate go run cmd/gen/main.go -spec api_spec.json -genfile=$GOFILE -mode=api

package baller

import (
	"net/url"
)

type CommonAllPlayers_Result struct {
}

func (c *Client) CommonAllPlayers(LeagueID string, Season string, IsOnlyCurrentSeason bool) (*CommonAllPlayers_Result, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonallplayers?"
		dest CommonAllPlayers_Result
	)

	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("Season", encodeString(Season))
	q.Set("IsOnlyCurrentSeason", encodeBool(IsOnlyCurrentSeason))

	if err := c.do(url+q.Encode(), &dest); err != nil {
		return nil, err
	}

	return &dest, nil
}

type Scoreboard_Result struct {
}

func (c *Client) Scoreboard(GameDate string, LeagueID string, DayOffset int) (*Scoreboard_Result, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboard?"
		dest Scoreboard_Result
	)

	q.Set("GameDate", encodeString(GameDate))
	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("DayOffset", encodeInt(DayOffset))

	if err := c.do(url+q.Encode(), &dest); err != nil {
		return nil, err
	}

	return &dest, nil
}

type ScoreboardV2_GameHeader struct {
	GAME_DATE_EST string
	GAME_SEQUENCE int
	GAME_ID       string
}

type ScoreboardV2_Result struct {
	GameHeader []ScoreboardV2_GameHeader
}

func (c *Client) ScoreboardV2(GameDate string, LeagueID string, DayOffset int) (*ScoreboardV2_Result, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboardv2?"
		dest ScoreboardV2_Result
	)

	q.Set("GameDate", encodeString(GameDate))
	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("DayOffset", encodeInt(DayOffset))

	if err := c.do(url+q.Encode(), &dest); err != nil {
		return nil, err
	}

	return &dest, nil
}
