package baller

import (
	"net/url"
)

type TeamGameLogOptions struct {
	TeamID     int
	LeagueID   string
	Season     string
	SeasonType string
}

type TeamGameLogResponse struct {
	TeamGameLog []TeamGameLogTeamGameLog
}

func (c *Client) TeamGameLog(options *TeamGameLogOptions) (*TeamGameLogResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "teamgamelog?"
		dest TeamGameLogResponse
		res  result
	)

	q.Set("TeamID", encodeInt(options.TeamID))
	q.Set("LeagueID", encodeString(options.LeagueID))
	q.Set("Season", encodeString(options.Season))
	q.Set("SeasonType", encodeString(options.SeasonType))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"TeamGameLog": &dest.TeamGameLog,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type TeamGameLogTeamGameLog struct {
	TeamID   int     `header:"Team_ID"`
	GameID   string  `header:"Game_ID"`
	GameDate string  `header:"GAME_DATE"`
	Matchup  string  `header:"MATCHUP"`
	Wl       string  `header:"WL"`
	Min      int     `header:"MIN"`
	Fgm      int     `header:"FGM"`
	Fga      int     `header:"FGA"`
	FgPct    float32 `header:"FG_PCT"`
	Fg3m     int     `header:"FG3M"`
	Fg3a     int     `header:"FG3A"`
	Fg3Pct   float32 `header:"FG3_PCT"`
	Ftm      int     `header:"FTM"`
	Fta      int     `header:"FTA"`
	FtPct    float32 `header:"FT_PCT"`
	Oreb     int     `header:"OREB"`
	Dreb     int     `header:"DREB"`
	Reb      int     `header:"REB"`
	Ast      int     `header:"AST"`
	Stl      int     `header:"STL"`
	Blk      int     `header:"BLK"`
	Tov      int     `header:"TOV"`
	Pf       int     `header:"PF"`
	Pts      int     `header:"PTS"`
}
