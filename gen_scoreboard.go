package baller

import (
	"net/url"
)

type ScoreboardOptions struct {
	DayOffset string
	GameDate  string
	LeagueID  string
}

type ScoreboardResponse struct {
	GameHeader             []ScoreboardGameHeader
	LineScore              []ScoreboardLineScore
	SeriesStandings        []ScoreboardSeriesStandings
	LastMeeting            []ScoreboardLastMeeting
	EastConfStandingsByDay []ScoreboardEastConfStandingsByDay
	WestConfStandingsByDay []ScoreboardWestConfStandingsByDay
	Available              []ScoreboardAvailable
}

func (c *Client) Scoreboard(options *ScoreboardOptions) (*ScoreboardResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboard?"
		dest ScoreboardResponse
		res  result
	)

	q.Set("DayOffset", encodeString(options.DayOffset))
	q.Set("GameDate", encodeString(options.GameDate))
	q.Set("LeagueID", encodeString(options.LeagueID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"GameHeader":             &dest.GameHeader,
		"LineScore":              &dest.LineScore,
		"SeriesStandings":        &dest.SeriesStandings,
		"LastMeeting":            &dest.LastMeeting,
		"EastConfStandingsByDay": &dest.EastConfStandingsByDay,
		"WestConfStandingsByDay": &dest.WestConfStandingsByDay,
		"Available":              &dest.Available,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type ScoreboardGameHeader struct {
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

type ScoreboardLineScore struct {
	GameDateEst      string  `header:"GAME_DATE_EST"`
	GameSequence     int     `header:"GAME_SEQUENCE"`
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCityName     string  `header:"TEAM_CITY_NAME"`
	TeamWinsLosses   string  `header:"TEAM_WINS_LOSSES"`
	PtsQtr1          int     `header:"PTS_QTR1"`
	PtsQtr2          int     `header:"PTS_QTR2"`
	PtsQtr3          int     `header:"PTS_QTR3"`
	PtsQtr4          int     `header:"PTS_QTR4"`
	PtsOt1           int     `header:"PTS_OT1"`
	PtsOt2           int     `header:"PTS_OT2"`
	PtsOt3           int     `header:"PTS_OT3"`
	PtsOt4           int     `header:"PTS_OT4"`
	PtsOt5           int     `header:"PTS_OT5"`
	PtsOt6           int     `header:"PTS_OT6"`
	PtsOt7           int     `header:"PTS_OT7"`
	PtsOt8           int     `header:"PTS_OT8"`
	PtsOt9           int     `header:"PTS_OT9"`
	PtsOt10          int     `header:"PTS_OT10"`
	Pts              int     `header:"PTS"`
	FgPct            float32 `header:"FG_PCT"`
	FtPct            float32 `header:"FT_PCT"`
	Fg3Pct           float32 `header:"FG3_PCT"`
	Ast              int     `header:"AST"`
	Reb              int     `header:"REB"`
	Tov              int     `header:"TOV"`
}

type ScoreboardSeriesStandings struct {
	GameID         string `header:"GAME_ID"`
	HomeTeamID     int    `header:"HOME_TEAM_ID"`
	VisitorTeamID  int    `header:"VISITOR_TEAM_ID"`
	GameDateEst    string `header:"GAME_DATE_EST"`
	HomeTeamWins   int    `header:"HOME_TEAM_WINS"`
	HomeTeamLosses int    `header:"HOME_TEAM_LOSSES"`
	SeriesLeader   string `header:"SERIES_LEADER"`
}

type ScoreboardLastMeeting struct {
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

type ScoreboardEastConfStandingsByDay struct {
	TeamID        int     `header:"TEAM_ID"`
	LeagueID      string  `header:"LEAGUE_ID"`
	SeasonID      string  `header:"SEASON_ID"`
	Standingsdate string  `header:"STANDINGSDATE"`
	Conference    string  `header:"CONFERENCE"`
	Team          string  `header:"TEAM"`
	G             int     `header:"G"`
	W             int     `header:"W"`
	L             int     `header:"L"`
	WPct          float32 `header:"W_PCT"`
	HomeRecord    string  `header:"HOME_RECORD"`
	RoadRecord    string  `header:"ROAD_RECORD"`
}

type ScoreboardWestConfStandingsByDay struct {
	TeamID        int     `header:"TEAM_ID"`
	LeagueID      string  `header:"LEAGUE_ID"`
	SeasonID      string  `header:"SEASON_ID"`
	Standingsdate string  `header:"STANDINGSDATE"`
	Conference    string  `header:"CONFERENCE"`
	Team          string  `header:"TEAM"`
	G             int     `header:"G"`
	W             int     `header:"W"`
	L             int     `header:"L"`
	WPct          float32 `header:"W_PCT"`
	HomeRecord    string  `header:"HOME_RECORD"`
	RoadRecord    string  `header:"ROAD_RECORD"`
}

type ScoreboardAvailable struct {
	GameID      string `header:"GAME_ID"`
	PtAvailable int    `header:"PT_AVAILABLE"`
}
