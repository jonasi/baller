package baller

import (
	"net/url"
)

type BoxscorePlayerTrackV2Options struct {
	GameID string
}

type BoxscorePlayerTrackV2Response struct {
	PlayerTrack     []BoxscorePlayerTrackV2PlayerTrack
	PlayerTrackTeam []BoxscorePlayerTrackV2PlayerTrackTeam
}

func (c *Client) BoxscorePlayerTrackV2(options *BoxscorePlayerTrackV2Options) (*BoxscorePlayerTrackV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscoreplayertrackv2?"
		dest BoxscorePlayerTrackV2Response
		res  result
	)

	q.Set("GameID", encodeString(options.GameID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"PlayerTrack":     &dest.PlayerTrack,
		"PlayerTrackTeam": &dest.PlayerTrackTeam,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BoxscorePlayerTrackV2PlayerTrack struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	PlayerID         int     `header:"PLAYER_ID"`
	PlayerName       string  `header:"PLAYER_NAME"`
	StartPosition    string  `header:"START_POSITION"`
	Comment          string  `header:"COMMENT"`
	Min              string  `header:"MIN"`
	Spd              float32 `header:"SPD"`
	Dist             float32 `header:"DIST"`
	Orbc             int     `header:"ORBC"`
	Drbc             int     `header:"DRBC"`
	Rbc              int     `header:"RBC"`
	Tchs             int     `header:"TCHS"`
	Sast             int     `header:"SAST"`
	Ftast            int     `header:"FTAST"`
	Pass             int     `header:"PASS"`
	Ast              int     `header:"AST"`
	Cfgm             int     `header:"CFGM"`
	Cfga             int     `header:"CFGA"`
	CfgPct           float32 `header:"CFG_PCT"`
	Ufgm             int     `header:"UFGM"`
	Ufga             int     `header:"UFGA"`
	UfgPct           float32 `header:"UFG_PCT"`
	FgPct            float32 `header:"FG_PCT"`
	Dfgm             int     `header:"DFGM"`
	Dfga             int     `header:"DFGA"`
	DfgPct           float32 `header:"DFG_PCT"`
}

type BoxscorePlayerTrackV2PlayerTrackTeam struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamNickname     string  `header:"TEAM_NICKNAME"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	Min              string  `header:"MIN"`
	Dist             float32 `header:"DIST"`
	Orbc             int     `header:"ORBC"`
	Drbc             int     `header:"DRBC"`
	Rbc              int     `header:"RBC"`
	Tchs             int     `header:"TCHS"`
	Sast             int     `header:"SAST"`
	Ftast            int     `header:"FTAST"`
	Pass             int     `header:"PASS"`
	Ast              int     `header:"AST"`
	Cfgm             int     `header:"CFGM"`
	Cfga             int     `header:"CFGA"`
	CfgPct           float32 `header:"CFG_PCT"`
	Ufgm             int     `header:"UFGM"`
	Ufga             int     `header:"UFGA"`
	UfgPct           float32 `header:"UFG_PCT"`
	FgPct            float32 `header:"FG_PCT"`
	Dfgm             int     `header:"DFGM"`
	Dfga             int     `header:"DFGA"`
	DfgPct           float32 `header:"DFG_PCT"`
}
