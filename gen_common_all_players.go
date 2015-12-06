package baller

import (
	"net/url"
)

type CommonAllPlayersOptions struct {
	LeagueID            string
	Season              string
	IsOnlyCurrentSeason int
}

type CommonAllPlayersResponse struct {
	CommonAllPlayers []CommonAllPlayersCommonAllPlayers
}

func (c *Client) CommonAllPlayers(options *CommonAllPlayersOptions) (*CommonAllPlayersResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonallplayers?"
		dest CommonAllPlayersResponse
		res  result
	)

	q.Set("LeagueID", encodeString(options.LeagueID))
	q.Set("Season", encodeString(options.Season))
	q.Set("IsOnlyCurrentSeason", encodeInt(options.IsOnlyCurrentSeason))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"CommonAllPlayers": &dest.CommonAllPlayers,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type CommonAllPlayersCommonAllPlayers struct {
	PersonID              int    `header:"PERSON_ID"`
	DisplayLastCommaFirst string `header:"DISPLAY_LAST_COMMA_FIRST"`
	Rosterstatus          int    `header:"ROSTERSTATUS"`
	FromYear              string `header:"FROM_YEAR"`
	ToYear                string `header:"TO_YEAR"`
	Playercode            string `header:"PLAYERCODE"`
	TeamID                int    `header:"TEAM_ID"`
	TeamCity              string `header:"TEAM_CITY"`
	TeamName              string `header:"TEAM_NAME"`
	TeamAbbreviation      string `header:"TEAM_ABBREVIATION"`
	TeamCode              string `header:"TEAM_CODE"`
	GamesPlayedFlag       string `header:"GAMES_PLAYED_FLAG"`
}
