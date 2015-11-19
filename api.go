//go:generate go run cmd/gen/main.go -spec api_spec.json -genfile=$GOFILE -mode=api

package baller

import (
	"net/url"
)

type ResultGameHeader struct {
	GAME_DATE_EST                    string `header:"GAME_DATE_EST"`
	GAME_SEQUENCE                    int    `header:"GAME_SEQUENCE"`
	GAME_ID                          string `header:"GAME_ID"`
	GAME_STATUS_ID                   int    `header:"GAME_STATUS_ID"`
	GAME_STATUS_TEXT                 string `header:"GAME_STATUS_TEXT"`
	GAMECODE                         string `header:"GAMECODE"`
	HOME_TEAM_ID                     int    `header:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID                  int    `header:"VISITOR_TEAM_ID"`
	SEASON                           string `header:"SEASON"`
	LIVE_PERIOD                      int    `header:"LIVE_PERIOD"`
	LIVE_PC_TIME                     string `header:"LIVE_PC_TIME"`
	NATL_TV_BROADCASTER_ABBREVIATION string `header:"NATL_TV_BROADCASTER_ABBREVIATION"`
	LIVE_PERIOD_TIME_BCAST           string `header:"LIVE_PERIOD_TIME_BCAST"`
	WH_STATUS                        int    `header:"WH_STATUS"`
}

type ResultLineScore struct {
	GAME_DATE_EST     string  `header:"GAME_DATE_EST"`
	GAME_SEQUENCE     int     `header:"GAME_SEQUENCE"`
	GAME_ID           string  `header:"GAME_ID"`
	TEAM_ID           int     `header:"TEAM_ID"`
	TEAM_ABBREVIATION string  `header:"TEAM_ABBREVIATION"`
	TEAM_CITY_NAME    string  `header:"TEAM_CITY_NAME"`
	TEAM_WINS_LOSSES  string  `header:"TEAM_WINS_LOSSES"`
	PTS_QTR1          int     `header:"PTS_QTR1"`
	PTS_QTR2          int     `header:"PTS_QTR2"`
	PTS_QTR3          int     `header:"PTS_QTR3"`
	PTS_QTR4          int     `header:"PTS_QTR4"`
	PTS_OT1           int     `header:"PTS_OT1"`
	PTS_OT2           int     `header:"PTS_OT2"`
	PTS_OT3           int     `header:"PTS_OT3"`
	PTS_OT4           int     `header:"PTS_OT4"`
	PTS_OT5           int     `header:"PTS_OT5"`
	PTS_OT6           int     `header:"PTS_OT6"`
	PTS_OT7           int     `header:"PTS_OT7"`
	PTS_OT8           int     `header:"PTS_OT8"`
	PTS_OT9           int     `header:"PTS_OT9"`
	PTS_OT10          int     `header:"PTS_OT10"`
	PTS               int     `header:"PTS"`
	FG_PCT            float32 `header:"FG_PCT"`
	FT_PCT            float32 `header:"FT_PCT"`
	FG3_PCT           float32 `header:"FG3_PCT"`
	AST               int     `header:"AST"`
	REB               int     `header:"REB"`
	TOV               int     `header:"TOV"`
}

type ResultSeriesStandings struct {
	GAME_ID          string `header:"GAME_ID"`
	HOME_TEAM_ID     int    `header:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID  int    `header:"VISITOR_TEAM_ID"`
	GAME_DATE_EST    string `header:"GAME_DATE_EST"`
	HOME_TEAM_WINS   int    `header:"HOME_TEAM_WINS"`
	HOME_TEAM_LOSSES int    `header:"HOME_TEAM_LOSSES"`
	SERIES_LEADER    string `header:"SERIES_LEADER"`
}

type ResultLastMeeting struct {
	GAME_ID                          string `header:"GAME_ID"`
	LAST_GAME_ID                     string `header:"LAST_GAME_ID"`
	LAST_GAME_DATE_EST               string `header:"LAST_GAME_DATE_EST"`
	LAST_GAME_HOME_TEAM_ID           int    `header:"LAST_GAME_HOME_TEAM_ID"`
	LAST_GAME_HOME_TEAM_CITY         string `header:"LAST_GAME_HOME_TEAM_CITY"`
	LAST_GAME_HOME_TEAM_NAME         string `header:"LAST_GAME_HOME_TEAM_NAME"`
	LAST_GAME_HOME_TEAM_ABBREVIATION string `header:"LAST_GAME_HOME_TEAM_ABBREVIATION"`
	LAST_GAME_HOME_TEAM_POINTS       int    `header:"LAST_GAME_HOME_TEAM_POINTS"`
	LAST_GAME_VISITOR_TEAM_ID        int    `header:"LAST_GAME_VISITOR_TEAM_ID"`
	LAST_GAME_VISITOR_TEAM_CITY      string `header:"LAST_GAME_VISITOR_TEAM_CITY"`
	LAST_GAME_VISITOR_TEAM_NAME      string `header:"LAST_GAME_VISITOR_TEAM_NAME"`
	LAST_GAME_VISITOR_TEAM_CITY1     string `header:"LAST_GAME_VISITOR_TEAM_CITY1"`
	LAST_GAME_VISITOR_TEAM_POINTS    int    `header:"LAST_GAME_VISITOR_TEAM_POINTS"`
}

type ResultConferenceStandings struct {
	TEAM_ID       int     `header:"TEAM_ID"`
	LEAGUE_ID     string  `header:"LEAGUE_ID"`
	SEASON_ID     string  `header:"SEASON_ID"`
	STANDINGSDATE string  `header:"STANDINGSDATE"`
	CONFERENCE    string  `header:"CONFERENCE"`
	TEAM          string  `header:"TEAM"`
	GAME          int     `header:"GAME"`
	W             int     `header:"W"`
	L             int     `header:"L"`
	W_PCT         float32 `header:"W_PCT"`
	HOME_RECORD   string  `header:"HOME_RECORD"`
	ROAD_RECORD   string  `header:"ROAD_RECORD"`
}

type ResultAvailable struct {
	GAME_ID      string `header:"GAME_ID"`
	PT_AVAILABLE int    `header:"PT_AVAILABLE"`
}

type ResultTeamLeaders struct {
	GAME_ID           string `header:"GAME_ID"`
	TEAM_ID           int    `header:"TEAM_ID"`
	TEAM_CITY         string `header:"TEAM_CITY"`
	TEAM_NICKNAME     string `header:"TEAM_NICKNAME"`
	TEAM_ABBREVIATION string `header:"TEAM_ABBREVIATION"`
	PTS_PLAYER_ID     int    `header:"PTS_PLAYER_ID"`
	PTS_PLAYER_NAME   string `header:"PTS_PLAYER_NAME"`
	PTS               int    `header:"PTS"`
	REB_PLAYER_ID     int    `header:"REB_PLAYER_ID"`
	REB_PLAYER_NAME   string `header:"REB_PLAYER_NAME"`
	REB               int    `header:"REB"`
	AST_PLAYER_ID     int    `header:"AST_PLAYER_ID"`
	AST_PLAYER_NAME   string `header:"AST_PLAYER_NAME"`
	AST               int    `header:"AST"`
}

type ResultTicketLinks struct {
	GAME_ID  string `header:"GAME_ID"`
	LEAG_TIX string `header:"LEAG_TIX"`
}

type ResultPlayer struct {
	PERSON_ID                int    `header:"PERSON_ID"`
	DISPLAY_LAST_COMMA_FIRST string `header:"DISPLAY_LAST_COMMA_FIRST"`
	ROSTERSTATUS             int    `header:"ROSTERSTATUS"`
	FROM_YEAR                string `header:"FROM_YEAR"`
	TO_YEAR                  string `header:"TO_YEAR"`
	PLAYERCODE               string `header:"PLAYERCODE"`
	TEAM_ID                  int    `header:"TEAM_ID"`
	TEAM_CITY                string `header:"TEAM_CITY"`
	TEAM_NAME                string `header:"TEAM_NAME"`
	TEAM_ABBREVIATION        string `header:"TEAM_ABBREVIATION"`
	TEAM_CODE                string `header:"TEAM_CODE"`
	GAMES_PLAYED_FLAG        string `header:"GAMES_PLAYED_FLAG"`
}

type CommonAllPlayers_Result struct {
	CommonAllPlayers []ResultPlayer
}

func (c *Client) CommonAllPlayers(LeagueID string, Season string, IsOnlyCurrentSeason bool) (*CommonAllPlayers_Result, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonallplayers?"
		dest CommonAllPlayers_Result
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

type Scoreboard_Result struct {
	GameHeader             []ResultGameHeader
	LineScore              []ResultLineScore
	SeriesStandings        []ResultSeriesStandings
	LastMeeting            []ResultLastMeeting
	EastConfStandingsByDay []ResultConferenceStandings
	WestConfStandingsByDay []ResultConferenceStandings
	Available              []ResultAvailable
}

func (c *Client) Scoreboard(GameDate string, LeagueID string, DayOffset int) (*Scoreboard_Result, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboard?"
		dest Scoreboard_Result
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

type ScoreboardV2_Result struct {
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

func (c *Client) ScoreboardV2(GameDate string, LeagueID string, DayOffset int) (*ScoreboardV2_Result, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "scoreboardv2?"
		dest ScoreboardV2_Result
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
