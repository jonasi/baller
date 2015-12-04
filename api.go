//go:generate go run cmd/gen/main.go -spec api_spec.json -genfile=$GOFILE -mode=api

package baller

import (
	"net/url"
)

type ResultGameHeader struct {
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

type ResultLineScore struct {
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

type ResultSeriesStandings struct {
	GameID         string `header:"GAME_ID"`
	HomeTeamID     int    `header:"HOME_TEAM_ID"`
	VisitorTeamID  int    `header:"VISITOR_TEAM_ID"`
	GameDateEst    string `header:"GAME_DATE_EST"`
	HomeTeamWins   int    `header:"HOME_TEAM_WINS"`
	HomeTeamLosses int    `header:"HOME_TEAM_LOSSES"`
	SeriesLeader   string `header:"SERIES_LEADER"`
}

type ResultLastMeeting struct {
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

type ResultConferenceStandings struct {
	TeamID        int     `header:"TEAM_ID"`
	LeagueID      string  `header:"LEAGUE_ID"`
	SeasonID      string  `header:"SEASON_ID"`
	Standingsdate string  `header:"STANDINGSDATE"`
	Conference    string  `header:"CONFERENCE"`
	Team          string  `header:"TEAM"`
	Game          int     `header:"GAME"`
	W             int     `header:"W"`
	L             int     `header:"L"`
	WPct          float32 `header:"W_PCT"`
	HomeRecord    string  `header:"HOME_RECORD"`
	RoadRecord    string  `header:"ROAD_RECORD"`
}

type ResultAvailable struct {
	GameID      string `header:"GAME_ID"`
	PtAvailable int    `header:"PT_AVAILABLE"`
}

type ResultTeamLeaders struct {
	GameID           string `header:"GAME_ID"`
	TeamID           int    `header:"TEAM_ID"`
	TeamCity         string `header:"TEAM_CITY"`
	TeamNickname     string `header:"TEAM_NICKNAME"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
	PtsPlayerID      int    `header:"PTS_PLAYER_ID"`
	PtsPlayerName    string `header:"PTS_PLAYER_NAME"`
	Pts              int    `header:"PTS"`
	RebPlayerID      int    `header:"REB_PLAYER_ID"`
	RebPlayerName    string `header:"REB_PLAYER_NAME"`
	Reb              int    `header:"REB"`
	AstPlayerID      int    `header:"AST_PLAYER_ID"`
	AstPlayerName    string `header:"AST_PLAYER_NAME"`
	Ast              int    `header:"AST"`
}

type ResultTicketLinks struct {
	GameID  string `header:"GAME_ID"`
	LeagTix string `header:"LEAG_TIX"`
}

type ResultPlayer struct {
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

type ResultPlayerInfo struct {
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

type ResultPlayerHeadlineStats struct {
	PlayerID   int     `header:"PLAYER_ID"`
	PlayerName string  `header:"PLAYER_NAME"`
	Timeframe  string  `header:"TimeFrame"`
	Pts        float32 `header:"PTS"`
	Ast        float32 `header:"AST"`
	Reb        float32 `header:"REB"`
	Pie        float32 `header:"PIE"`
}

type CommonPlayerInfoResponse struct {
	CommonPlayerInfo    []ResultPlayerInfo
	PlayerHeadlineStats []ResultPlayerHeadlineStats
}

func (c *Client) CommonPlayerInfo(PlayerID int, LeagueID string) (*CommonPlayerInfoResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonplayerinfo?"
		dest CommonPlayerInfoResponse
		res  result
	)

	q.Set("PlayerID", encodeInt(PlayerID))
	q.Set("LeagueID", encodeString(LeagueID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("CommonPlayerInfo", ResultPlayerInfo{}); err == nil {
		dest.CommonPlayerInfo = d.([]ResultPlayerInfo)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("PlayerHeadlineStats", ResultPlayerHeadlineStats{}); err == nil {
		dest.PlayerHeadlineStats = d.([]ResultPlayerHeadlineStats)
	} else {
		return nil, err
	}

	return &dest, nil
}

type CommonAllPlayersResponse struct {
	CommonAllPlayers []ResultPlayer
}

func (c *Client) CommonAllPlayers(LeagueID string, Season string, IsOnlyCurrentSeason bool) (*CommonAllPlayersResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonallplayers?"
		dest CommonAllPlayersResponse
		res  result
	)

	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("Season", encodeString(Season))
	q.Set("IsOnlyCurrentSeason", encodeBool(IsOnlyCurrentSeason))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("CommonAllPlayers", ResultPlayer{}); err == nil {
		dest.CommonAllPlayers = d.([]ResultPlayer)
	} else {
		return nil, err
	}

	return &dest, nil
}

type ScoreboardResponse struct {
	GameHeader             []ResultGameHeader
	LineScore              []ResultLineScore
	SeriesStandings        []ResultSeriesStandings
	LastMeeting            []ResultLastMeeting
	EastConfStandingsByDay []ResultConferenceStandings
	WestConfStandingsByDay []ResultConferenceStandings
	Available              []ResultAvailable
}

func (c *Client) Scoreboard(GameDate string, LeagueID string, DayOffset int) (*ScoreboardResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboard?"
		dest ScoreboardResponse
		res  result
	)

	q.Set("GameDate", encodeString(GameDate))
	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("DayOffset", encodeInt(DayOffset))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("GameHeader", ResultGameHeader{}); err == nil {
		dest.GameHeader = d.([]ResultGameHeader)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("LineScore", ResultLineScore{}); err == nil {
		dest.LineScore = d.([]ResultLineScore)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("SeriesStandings", ResultSeriesStandings{}); err == nil {
		dest.SeriesStandings = d.([]ResultSeriesStandings)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("LastMeeting", ResultLastMeeting{}); err == nil {
		dest.LastMeeting = d.([]ResultLastMeeting)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("EastConfStandingsByDay", ResultConferenceStandings{}); err == nil {
		dest.EastConfStandingsByDay = d.([]ResultConferenceStandings)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("WestConfStandingsByDay", ResultConferenceStandings{}); err == nil {
		dest.WestConfStandingsByDay = d.([]ResultConferenceStandings)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("Available", ResultAvailable{}); err == nil {
		dest.Available = d.([]ResultAvailable)
	} else {
		return nil, err
	}

	return &dest, nil
}

type ScoreboardV2Response struct {
	GameHeader             []ResultGameHeader
	LineScore              []ResultLineScore
	SeriesStandings        []ResultSeriesStandings
	LastMeeting            []ResultLastMeeting
	EastConfStandingsByDay []ResultConferenceStandings
	WestConfStandingsByDay []ResultConferenceStandings
	Available              []ResultAvailable
	TeamLeaders            []ResultTeamLeaders
	TicketLinks            []ResultTicketLinks
}

func (c *Client) ScoreboardV2(GameDate string, LeagueID string, DayOffset int) (*ScoreboardV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboardv2?"
		dest ScoreboardV2Response
		res  result
	)

	q.Set("GameDate", encodeString(GameDate))
	q.Set("LeagueID", encodeString(LeagueID))
	q.Set("DayOffset", encodeInt(DayOffset))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("GameHeader", ResultGameHeader{}); err == nil {
		dest.GameHeader = d.([]ResultGameHeader)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("LineScore", ResultLineScore{}); err == nil {
		dest.LineScore = d.([]ResultLineScore)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("SeriesStandings", ResultSeriesStandings{}); err == nil {
		dest.SeriesStandings = d.([]ResultSeriesStandings)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("LastMeeting", ResultLastMeeting{}); err == nil {
		dest.LastMeeting = d.([]ResultLastMeeting)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("EastConfStandingsByDay", ResultConferenceStandings{}); err == nil {
		dest.EastConfStandingsByDay = d.([]ResultConferenceStandings)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("WestConfStandingsByDay", ResultConferenceStandings{}); err == nil {
		dest.WestConfStandingsByDay = d.([]ResultConferenceStandings)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("Available", ResultAvailable{}); err == nil {
		dest.Available = d.([]ResultAvailable)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("TeamLeaders", ResultTeamLeaders{}); err == nil {
		dest.TeamLeaders = d.([]ResultTeamLeaders)
	} else {
		return nil, err
	}

	if d, err := res.unmarshalResultSet("TicketLinks", ResultTicketLinks{}); err == nil {
		dest.TicketLinks = d.([]ResultTicketLinks)
	} else {
		return nil, err
	}

	return &dest, nil
}
