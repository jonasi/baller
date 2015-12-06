package baller

import (
	"net/url"
)

type BoxscoreMiscV2SqlPlayersMisc struct {
	GameID           string `header:"GAME_ID"`
	TeamID           int    `header:"TEAM_ID"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
	TeamCity         string `header:"TEAM_CITY"`
	PlayerID         int    `header:"PLAYER_ID"`
	PlayerName       string `header:"PLAYER_NAME"`
	StartPosition    string `header:"START_POSITION"`
	Comment          string `header:"COMMENT"`
	Min              string `header:"MIN"`
	PtsOffTov        int    `header:"PTS_OFF_TOV"`
	Pts2ndChance     int    `header:"PTS_2ND_CHANCE"`
	PtsFb            int    `header:"PTS_FB"`
	PtsPaint         int    `header:"PTS_PAINT"`
	OppPtsOffTov     int    `header:"OPP_PTS_OFF_TOV"`
	OppPts2ndChance  int    `header:"OPP_PTS_2ND_CHANCE"`
	OppPtsFb         int    `header:"OPP_PTS_FB"`
	OppPtsPaint      int    `header:"OPP_PTS_PAINT"`
	Blk              int    `header:"BLK"`
	Blka             int    `header:"BLKA"`
	Pf               int    `header:"PF"`
	Pfd              int    `header:"PFD"`
}

type BoxscoreMiscV2SqlTeamsMisc struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamName         string  `header:"TEAM_NAME"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	Min              string  `header:"MIN"`
	PtsOffTov        float32 `header:"PTS_OFF_TOV"`
	Pts2ndChance     float32 `header:"PTS_2ND_CHANCE"`
	PtsFb            float32 `header:"PTS_FB"`
	PtsPaint         float32 `header:"PTS_PAINT"`
	OppPtsOffTov     float32 `header:"OPP_PTS_OFF_TOV"`
	OppPts2ndChance  float32 `header:"OPP_PTS_2ND_CHANCE"`
	OppPtsFb         float32 `header:"OPP_PTS_FB"`
	OppPtsPaint      float32 `header:"OPP_PTS_PAINT"`
	Blk              int     `header:"BLK"`
	Blka             int     `header:"BLKA"`
	Pf               int     `header:"PF"`
	Pfd              int     `header:"PFD"`
}

type BoxscoreMiscV2Response struct {
	SqlPlayersMisc []BoxscoreMiscV2SqlPlayersMisc
	SqlTeamsMisc   []BoxscoreMiscV2SqlTeamsMisc
}

type BoxscoreMiscV2Options struct {
	StartRange  int
	EndRange    int
	RangeType   int
	GameID      string
	StartPeriod int
	EndPeriod   int
}

func (c *Client) BoxscoreMiscV2(options *BoxscoreMiscV2Options) (*BoxscoreMiscV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscore_misc_v2?"
		dest BoxscoreMiscV2Response
		res  result
	)

	q.Set("StartRange", encodeInt(options.StartRange))
	q.Set("EndRange", encodeInt(options.EndRange))
	q.Set("RangeType", encodeInt(options.RangeType))
	q.Set("GameID", encodeString(options.GameID))
	q.Set("StartPeriod", encodeInt(options.StartPeriod))
	q.Set("EndPeriod", encodeInt(options.EndPeriod))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"sqlPlayersMisc": &dest.SqlPlayersMisc,
		"sqlTeamsMisc":   &dest.SqlTeamsMisc,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}
