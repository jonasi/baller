package baller

import (
	"net/url"
)

type CommonPlayerInfoOptions struct {
	LeagueID string
	PlayerID int
}

type CommonPlayerInfoResponse struct {
	CommonPlayerInfo    []CommonPlayerInfoCommonPlayerInfo
	PlayerHeadlineStats []CommonPlayerInfoPlayerHeadlineStats
}

func (c *Client) CommonPlayerInfo(options *CommonPlayerInfoOptions) (*CommonPlayerInfoResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonplayerinfo?"
		dest CommonPlayerInfoResponse
		res  result
	)

	q.Set("LeagueID", encodeString(options.LeagueID))
	q.Set("PlayerID", encodeInt(options.PlayerID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"CommonPlayerInfo":    &dest.CommonPlayerInfo,
		"PlayerHeadlineStats": &dest.PlayerHeadlineStats,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type CommonPlayerInfoCommonPlayerInfo struct {
	PersonID              int    `header:"PERSON_ID"`
	FirstName             string `header:"FIRST_NAME"`
	LastName              string `header:"LAST_NAME"`
	DisplayFirstLast      string `header:"DISPLAY_FIRST_LAST"`
	DisplayLastCommaFirst string `header:"DISPLAY_LAST_COMMA_FIRST"`
	DisplayFiLast         string `header:"DISPLAY_FI_LAST"`
	Birthdate             string `header:"BIRTHDATE"`
	School                string `header:"SCHOOL"`
	Country               string `header:"COUNTRY"`
	LastAffiliation       string `header:"LAST_AFFILIATION"`
	Height                string `header:"HEIGHT"`
	Weight                string `header:"WEIGHT"`
	SeasonExp             int    `header:"SEASON_EXP"`
	Jersey                string `header:"JERSEY"`
	Position              string `header:"POSITION"`
	Rosterstatus          string `header:"ROSTERSTATUS"`
	TeamID                int    `header:"TEAM_ID"`
	TeamName              string `header:"TEAM_NAME"`
	TeamAbbreviation      string `header:"TEAM_ABBREVIATION"`
	TeamCode              string `header:"TEAM_CODE"`
	TeamCity              string `header:"TEAM_CITY"`
	Playercode            string `header:"PLAYERCODE"`
	FromYear              int    `header:"FROM_YEAR"`
	ToYear                int    `header:"TO_YEAR"`
	DleagueFlag           string `header:"DLEAGUE_FLAG"`
	GamesPlayedFlag       string `header:"GAMES_PLAYED_FLAG"`
}

type CommonPlayerInfoPlayerHeadlineStats struct {
	PlayerID   int     `header:"PLAYER_ID"`
	PlayerName string  `header:"PLAYER_NAME"`
	Timeframe  string  `header:"TimeFrame"`
	Pts        float32 `header:"PTS"`
	Ast        float32 `header:"AST"`
	Reb        float32 `header:"REB"`
	Pie        float32 `header:"PIE"`
}
