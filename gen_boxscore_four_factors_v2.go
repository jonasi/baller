package baller

import (
	"net/url"
)

type BoxscoreFourFactorsV2Options struct {
	GameID      string
	StartPeriod int
	EndPeriod   int
	StartRange  int
	EndRange    int
	RangeType   int
}

type BoxscoreFourFactorsV2Response struct {
	SqlPlayersFourFactors []BoxscoreFourFactorsV2SqlPlayersFourFactors
	SqlTeamsFourFactors   []BoxscoreFourFactorsV2SqlTeamsFourFactors
}

func (c *Client) BoxscoreFourFactorsV2(options *BoxscoreFourFactorsV2Options) (*BoxscoreFourFactorsV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscorefourfactorsv2?"
		dest BoxscoreFourFactorsV2Response
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
		"sqlPlayersFourFactors": &dest.SqlPlayersFourFactors,
		"sqlTeamsFourFactors":   &dest.SqlTeamsFourFactors,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BoxscoreFourFactorsV2SqlPlayersFourFactors struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	PlayerID         int     `header:"PLAYER_ID"`
	PlayerName       string  `header:"PLAYER_NAME"`
	StartPosition    string  `header:"START_POSITION"`
	Comment          string  `header:"COMMENT"`
	Min              string  `header:"MIN"`
	EfgPct           float32 `header:"EFG_PCT"`
	FtaRate          float32 `header:"FTA_RATE"`
	TmTovPct         float32 `header:"TM_TOV_PCT"`
	OrebPct          float32 `header:"OREB_PCT"`
	OppEfgPct        float32 `header:"OPP_EFG_PCT"`
	OppFtaRate       float32 `header:"OPP_FTA_RATE"`
	OppTovPct        float32 `header:"OPP_TOV_PCT"`
	OppOrebPct       float32 `header:"OPP_OREB_PCT"`
}

type BoxscoreFourFactorsV2SqlTeamsFourFactors struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamName         string  `header:"TEAM_NAME"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	Min              string  `header:"MIN"`
	EfgPct           float32 `header:"EFG_PCT"`
	FtaRate          float32 `header:"FTA_RATE"`
	TmTovPct         float32 `header:"TM_TOV_PCT"`
	OrebPct          float32 `header:"OREB_PCT"`
	OppEfgPct        float32 `header:"OPP_EFG_PCT"`
	OppFtaRate       float32 `header:"OPP_FTA_RATE"`
	OppTovPct        float32 `header:"OPP_TOV_PCT"`
	OppOrebPct       float32 `header:"OPP_OREB_PCT"`
}
