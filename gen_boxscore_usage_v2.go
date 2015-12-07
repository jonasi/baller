package baller

import (
	"net/url"
)

type BoxscoreUsageV2Options struct {
	StartPeriod int
	EndPeriod   int
	StartRange  int
	EndRange    int
	RangeType   int
	GameID      string
}

type BoxscoreUsageV2Response struct {
	SqlPlayersUsage []BoxscoreUsageV2SqlPlayersUsage
	SqlTeamsUsage   []BoxscoreUsageV2SqlTeamsUsage
}

func (c *Client) BoxscoreUsageV2(options *BoxscoreUsageV2Options) (*BoxscoreUsageV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscoreusagev2?"
		dest BoxscoreUsageV2Response
		res  result
	)

	q.Set("StartPeriod", encodeInt(options.StartPeriod))
	q.Set("EndPeriod", encodeInt(options.EndPeriod))
	q.Set("StartRange", encodeInt(options.StartRange))
	q.Set("EndRange", encodeInt(options.EndRange))
	q.Set("RangeType", encodeInt(options.RangeType))
	q.Set("GameID", encodeString(options.GameID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"sqlPlayersUsage": &dest.SqlPlayersUsage,
		"sqlTeamsUsage":   &dest.SqlTeamsUsage,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BoxscoreUsageV2SqlPlayersUsage struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	PlayerID         int     `header:"PLAYER_ID"`
	PlayerName       string  `header:"PLAYER_NAME"`
	StartPosition    string  `header:"START_POSITION"`
	Comment          string  `header:"COMMENT"`
	Min              string  `header:"MIN"`
	UsgPct           float32 `header:"USG_PCT"`
	PctFgm           float32 `header:"PCT_FGM"`
	PctFga           float32 `header:"PCT_FGA"`
	PctFg3m          float32 `header:"PCT_FG3M"`
	PctFg3a          float32 `header:"PCT_FG3A"`
	PctFtm           float32 `header:"PCT_FTM"`
	PctFta           float32 `header:"PCT_FTA"`
	PctOreb          float32 `header:"PCT_OREB"`
	PctDreb          float32 `header:"PCT_DREB"`
	PctReb           float32 `header:"PCT_REB"`
	PctAst           float32 `header:"PCT_AST"`
	PctTov           float32 `header:"PCT_TOV"`
	PctStl           float32 `header:"PCT_STL"`
	PctBlk           float32 `header:"PCT_BLK"`
	PctBlka          float32 `header:"PCT_BLKA"`
	PctPf            float32 `header:"PCT_PF"`
	PctPfd           float32 `header:"PCT_PFD"`
	PctPts           float32 `header:"PCT_PTS"`
}

type BoxscoreUsageV2SqlTeamsUsage struct {
	GameID           string `header:"GAME_ID"`
	TeamID           int    `header:"TEAM_ID"`
	TeamName         string `header:"TEAM_NAME"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
	TeamCity         string `header:"TEAM_CITY"`
	Min              string `header:"MIN"`
	UsgPct           int    `header:"USG_PCT"`
	PctFgm           int    `header:"PCT_FGM"`
	PctFga           int    `header:"PCT_FGA"`
	PctFg3m          int    `header:"PCT_FG3M"`
	PctFg3a          int    `header:"PCT_FG3A"`
	PctFtm           int    `header:"PCT_FTM"`
	PctFta           int    `header:"PCT_FTA"`
	PctOreb          int    `header:"PCT_OREB"`
	PctDreb          int    `header:"PCT_DREB"`
	PctReb           int    `header:"PCT_REB"`
	PctAst           int    `header:"PCT_AST"`
	PctTov           int    `header:"PCT_TOV"`
	PctStl           int    `header:"PCT_STL"`
	PctBlk           int    `header:"PCT_BLK"`
	PctBlka          int    `header:"PCT_BLKA"`
	PctPf            int    `header:"PCT_PF"`
	PctPfd           int    `header:"PCT_PFD"`
	PctPts           int    `header:"PCT_PTS"`
}
