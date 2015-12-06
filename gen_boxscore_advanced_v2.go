package baller

import (
	"net/url"
)

type BoxscoreAdvancedV2Options struct {
	GameID      string
	StartPeriod int
	EndPeriod   int
	StartRange  int
	EndRange    int
	RangeType   int
}

type BoxscoreAdvancedV2Response struct {
	PlayerStats []BoxscoreAdvancedV2PlayerStats
	TeamStats   []BoxscoreAdvancedV2TeamStats
}

func (c *Client) BoxscoreAdvancedV2(options *BoxscoreAdvancedV2Options) (*BoxscoreAdvancedV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscoreadvancedv2?"
		dest BoxscoreAdvancedV2Response
		res  result
	)

	q.Set("GameID", encodeString(options.GameID))
	q.Set("StartPeriod", encodeInt(options.StartPeriod))
	q.Set("EndPeriod", encodeInt(options.EndPeriod))
	q.Set("StartRange", encodeInt(options.StartRange))
	q.Set("EndRange", encodeInt(options.EndRange))
	q.Set("RangeType", encodeInt(options.RangeType))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"PlayerStats": &dest.PlayerStats,
		"TeamStats":   &dest.TeamStats,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BoxscoreAdvancedV2PlayerStats struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	PlayerID         int     `header:"PLAYER_ID"`
	PlayerName       string  `header:"PLAYER_NAME"`
	StartPosition    string  `header:"START_POSITION"`
	Comment          string  `header:"COMMENT"`
	Min              string  `header:"MIN"`
	OffRating        float32 `header:"OFF_RATING"`
	DefRating        float32 `header:"DEF_RATING"`
	NetRating        float32 `header:"NET_RATING"`
	AstPct           float32 `header:"AST_PCT"`
	AstTov           float32 `header:"AST_TOV"`
	AstRatio         float32 `header:"AST_RATIO"`
	OrebPct          float32 `header:"OREB_PCT"`
	DrebPct          float32 `header:"DREB_PCT"`
	RebPct           float32 `header:"REB_PCT"`
	TmTovPct         float32 `header:"TM_TOV_PCT"`
	EfgPct           float32 `header:"EFG_PCT"`
	TsPct            float32 `header:"TS_PCT"`
	UsgPct           float32 `header:"USG_PCT"`
	Pace             float32 `header:"PACE"`
	Pie              float32 `header:"PIE"`
}

type BoxscoreAdvancedV2TeamStats struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamName         string  `header:"TEAM_NAME"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	Min              string  `header:"MIN"`
	OffRating        float32 `header:"OFF_RATING"`
	DefRating        float32 `header:"DEF_RATING"`
	NetRating        float32 `header:"NET_RATING"`
	AstPct           float32 `header:"AST_PCT"`
	AstTov           float32 `header:"AST_TOV"`
	AstRatio         float32 `header:"AST_RATIO"`
	OrebPct          float32 `header:"OREB_PCT"`
	DrebPct          float32 `header:"DREB_PCT"`
	RebPct           float32 `header:"REB_PCT"`
	TmTovPct         float32 `header:"TM_TOV_PCT"`
	EfgPct           float32 `header:"EFG_PCT"`
	TsPct            float32 `header:"TS_PCT"`
	UsgPct           float32 `header:"USG_PCT"`
	Pace             float32 `header:"PACE"`
	Pie              float32 `header:"PIE"`
}
