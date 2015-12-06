package baller

import (
	"net/url"
)

type BoxscoreOptions struct {
	GameID      string
	StartPeriod int
	EndPeriod   int
	StartRange  int
	EndRange    int
	RangeType   int
}

type BoxscoreResponse struct {
	GameSummary     []BoxscoreGameSummary
	LineScore       []BoxscoreLineScore
	SeasonSeries    []BoxscoreSeasonSeries
	LastMeeting     []BoxscoreLastMeeting
	PlayerStats     []BoxscorePlayerStats
	TeamStats       []BoxscoreTeamStats
	OtherStats      []BoxscoreOtherStats
	Officials       []BoxscoreOfficials
	GameInfo        []BoxscoreGameInfo
	InactivePlayers []BoxscoreInactivePlayers
	AvailableVideo  []BoxscoreAvailableVideo
	PlayerTrack     []BoxscorePlayerTrack
	PlayerTrackTeam []BoxscorePlayerTrackTeam
}

func (c *Client) Boxscore(options *BoxscoreOptions) (*BoxscoreResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscore?"
		dest BoxscoreResponse
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
		"GameSummary":     &dest.GameSummary,
		"LineScore":       &dest.LineScore,
		"SeasonSeries":    &dest.SeasonSeries,
		"LastMeeting":     &dest.LastMeeting,
		"PlayerStats":     &dest.PlayerStats,
		"TeamStats":       &dest.TeamStats,
		"OtherStats":      &dest.OtherStats,
		"Officials":       &dest.Officials,
		"GameInfo":        &dest.GameInfo,
		"InactivePlayers": &dest.InactivePlayers,
		"AvailableVideo":  &dest.AvailableVideo,
		"PlayerTrack":     &dest.PlayerTrack,
		"PlayerTrackTeam": &dest.PlayerTrackTeam,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BoxscoreGameSummary struct {
	GameDateEst                   string `header:"GAME_DATE_EST"`
	GameSequence                  int    `header:"GAME_SEQUENCE"`
	GameID                        string `header:"GAME_ID"`
	GameStatusID                  int    `header:"GAME_STATUS_ID"`
	GameStatusText                string `header:"GAME_STATUS_TEXT"`
	Gamecode                      string `header:"GAMECODE"`
	HomeTeamID                    int    `header:"HOME_TEAM_ID"`
	VisitorTeamID                 int    `header:"VISITOR_TEAM_ID"`
	Season                        string `header:"SEASON"`
	LivePeriod                    int    `header:"LIVE_PERIOD"`
	LivePcTime                    string `header:"LIVE_PC_TIME"`
	NatlTvBroadcasterAbbreviation string `header:"NATL_TV_BROADCASTER_ABBREVIATION"`
	LivePeriodTimeBcast           string `header:"LIVE_PERIOD_TIME_BCAST"`
	WhStatus                      int    `header:"WH_STATUS"`
}

type BoxscoreLineScore struct {
	GameDateEst      string `header:"GAME_DATE_EST"`
	GameSequence     int    `header:"GAME_SEQUENCE"`
	GameID           string `header:"GAME_ID"`
	TeamID           int    `header:"TEAM_ID"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
	TeamCityName     string `header:"TEAM_CITY_NAME"`
	TeamWinsLosses   string `header:"TEAM_WINS_LOSSES"`
	PtsQtr1          int    `header:"PTS_QTR1"`
	PtsQtr2          int    `header:"PTS_QTR2"`
	PtsQtr3          int    `header:"PTS_QTR3"`
	PtsQtr4          int    `header:"PTS_QTR4"`
	PtsOt1           int    `header:"PTS_OT1"`
	PtsOt2           int    `header:"PTS_OT2"`
	PtsOt3           int    `header:"PTS_OT3"`
	PtsOt4           int    `header:"PTS_OT4"`
	PtsOt5           int    `header:"PTS_OT5"`
	PtsOt6           int    `header:"PTS_OT6"`
	PtsOt7           int    `header:"PTS_OT7"`
	PtsOt8           int    `header:"PTS_OT8"`
	PtsOt9           int    `header:"PTS_OT9"`
	PtsOt10          int    `header:"PTS_OT10"`
	Pts              int    `header:"PTS"`
}

type BoxscoreSeasonSeries struct {
	GameID         string `header:"GAME_ID"`
	HomeTeamID     int    `header:"HOME_TEAM_ID"`
	VisitorTeamID  int    `header:"VISITOR_TEAM_ID"`
	GameDateEst    string `header:"GAME_DATE_EST"`
	HomeTeamWins   int    `header:"HOME_TEAM_WINS"`
	HomeTeamLosses int    `header:"HOME_TEAM_LOSSES"`
	SeriesLeader   string `header:"SERIES_LEADER"`
}

type BoxscoreLastMeeting struct {
	GameID                       string `header:"GAME_ID"`
	LastGameID                   string `header:"LAST_GAME_ID"`
	LastGameDateEst              string `header:"LAST_GAME_DATE_EST"`
	LastGameHomeTeamID           int    `header:"LAST_GAME_HOME_TEAM_ID"`
	LastGameHomeTeamCity         string `header:"LAST_GAME_HOME_TEAM_CITY"`
	LastGameHomeTeamName         string `header:"LAST_GAME_HOME_TEAM_NAME"`
	LastGameHomeTeamAbbreviation string `header:"LAST_GAME_HOME_TEAM_ABBREVIATION"`
	LastGameHomeTeamPoints       int    `header:"LAST_GAME_HOME_TEAM_POINTS"`
	LastGameVisitorTeamID        int    `header:"LAST_GAME_VISITOR_TEAM_ID"`
	LastGameVisitorTeamCity      string `header:"LAST_GAME_VISITOR_TEAM_CITY"`
	LastGameVisitorTeamName      string `header:"LAST_GAME_VISITOR_TEAM_NAME"`
	LastGameVisitorTeamCity1     string `header:"LAST_GAME_VISITOR_TEAM_CITY1"`
	LastGameVisitorTeamPoints    int    `header:"LAST_GAME_VISITOR_TEAM_POINTS"`
}

type BoxscorePlayerStats struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	PlayerID         int     `header:"PLAYER_ID"`
	PlayerName       string  `header:"PLAYER_NAME"`
	StartPosition    string  `header:"START_POSITION"`
	Comment          string  `header:"COMMENT"`
	Min              string  `header:"MIN"`
	Fgm              int     `header:"FGM"`
	Fga              int     `header:"FGA"`
	FgPct            float32 `header:"FG_PCT"`
	Fg3m             int     `header:"FG3M"`
	Fg3a             int     `header:"FG3A"`
	Fg3Pct           float32 `header:"FG3_PCT"`
	Ftm              int     `header:"FTM"`
	Fta              int     `header:"FTA"`
	FtPct            float32 `header:"FT_PCT"`
	Oreb             int     `header:"OREB"`
	Dreb             int     `header:"DREB"`
	Reb              int     `header:"REB"`
	Ast              int     `header:"AST"`
	Stl              int     `header:"STL"`
	Blk              int     `header:"BLK"`
	To               int     `header:"TO"`
	Pf               int     `header:"PF"`
	Pts              int     `header:"PTS"`
	PlusMinus        float32 `header:"PLUS_MINUS"`
}

type BoxscoreTeamStats struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamName         string  `header:"TEAM_NAME"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	Min              string  `header:"MIN"`
	Fgm              int     `header:"FGM"`
	Fga              int     `header:"FGA"`
	FgPct            float32 `header:"FG_PCT"`
	Fg3m             int     `header:"FG3M"`
	Fg3a             int     `header:"FG3A"`
	Fg3Pct           float32 `header:"FG3_PCT"`
	Ftm              int     `header:"FTM"`
	Fta              int     `header:"FTA"`
	FtPct            float32 `header:"FT_PCT"`
	Oreb             int     `header:"OREB"`
	Dreb             int     `header:"DREB"`
	Reb              int     `header:"REB"`
	Ast              int     `header:"AST"`
	Stl              int     `header:"STL"`
	Blk              int     `header:"BLK"`
	To               int     `header:"TO"`
	Pf               int     `header:"PF"`
	Pts              int     `header:"PTS"`
	PlusMinus        float32 `header:"PLUS_MINUS"`
}

type BoxscoreOtherStats struct {
	LeagueID         string `header:"LEAGUE_ID"`
	SeasonID         string `header:"SEASON_ID"`
	TeamID           int    `header:"TEAM_ID"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
	TeamCity         string `header:"TEAM_CITY"`
	PtsPaint         int    `header:"PTS_PAINT"`
	Pts2ndChance     int    `header:"PTS_2ND_CHANCE"`
	PtsFb            int    `header:"PTS_FB"`
	LargestLead      int    `header:"LARGEST_LEAD"`
	LeadChanges      int    `header:"LEAD_CHANGES"`
	TimesTied        int    `header:"TIMES_TIED"`
}

type BoxscoreOfficials struct {
	OfficialID int    `header:"OFFICIAL_ID"`
	FirstName  string `header:"FIRST_NAME"`
	LastName   string `header:"LAST_NAME"`
	JerseyNum  string `header:"JERSEY_NUM"`
}

type BoxscoreGameInfo struct {
	GameDate   string `header:"GAME_DATE"`
	Attendance int    `header:"ATTENDANCE"`
	GameTime   string `header:"GAME_TIME"`
}

type BoxscoreInactivePlayers struct {
	PlayerID         int    `header:"PLAYER_ID"`
	FirstName        string `header:"FIRST_NAME"`
	LastName         string `header:"LAST_NAME"`
	JerseyNum        string `header:"JERSEY_NUM"`
	TeamID           int    `header:"TEAM_ID"`
	TeamCity         string `header:"TEAM_CITY"`
	TeamName         string `header:"TEAM_NAME"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
}

type BoxscoreAvailableVideo struct {
	GameID             string `header:"GAME_ID"`
	VideoAvailableFlag int    `header:"VIDEO_AVAILABLE_FLAG"`
	PtAvailable        int    `header:"PT_AVAILABLE"`
}

type BoxscorePlayerTrack struct {
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

type BoxscorePlayerTrackTeam struct {
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
