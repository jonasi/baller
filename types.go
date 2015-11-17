package baller

import (
	"encoding/json"
	"reflect"
)

type scoreboard struct {
	Resource   string
	Parameters struct {
		GameDate  string
		LeagueID  string
		DayOffset string
	}
	ResultSets []struct {
		Name    string
		Headers []string
		RowSet  []json.RawMessage
	}
}

type Scoreboard struct {
	Date  string
	Games []Game
}

type Game struct {
	Header    GameHeader
	LineScore LineScore
}

// GAME_DATE_EST GAME_SEQUENCE GAME_ID GAME_STATUS_ID GAME_STATUS_TEXT GAMECODE HOME_TEAM_ID VISITOR_TEAM_ID SEASON LIVE_PERIOD LIVE_PC_TIME NATL_TV_BROADCASTER_ABBREVIATION LIVE_PERIOD_TIME_BCAST WH_STATUS
type GameHeader struct {
	GameDateEST             string
	GameSequence            int
	GameID                  string
	GameStatusID            int
	GameStatusText          string
	GameCode                string
	HomeTeamID              int
	VisitorTeamID           int
	Season                  string
	LivePeriod              int
	LivePCTime              string
	NatlTVBroadcasterAbbrev string
	LivePeriodTimeBcast     string
	WHStatus                int
}

// "GAME_DATE_EST", "GAME_SEQUENCE","GAME_ID","TEAM_ID","TEAM_ABBREVIATION","TEAM_CITY_NAME","TEAM_WINS_LOSSES","PTS_QTR1","PTS_QTR2","PTS_QTR3","PTS_QTR4","PTS_OT1","PTS_OT2","PTS_OT3","PTS_OT4","PTS_OT5","PTS_OT6","PTS_OT7 ","PTS_OT8","PTS_OT9","PTS_OT10","PTS","FG_PCT","FT_PCT","FG3_PCT","AST","REB","TOV"
type LineScore struct {
	GameDateEST    string
	GameSequence   int
	GameID         string
	TeamID         int
	TeamAbbrev     string
	TeamCityName   string
	TeamWinsLosses string
	PtsQtr1        int
	PtsQtr2        int
	PtsQtr3        int
	PtsQtr4        int
	PtsOT1         int
	PtsOT2         int
	PtsOT3         int
	PtsOT4         int
	PtsOT5         int
	PtsOT6         int
	PtsOT7         int
	PtsOT8         int
	PtsOT9         int
	PtsOT10        int
	Pts            int
	FGPct          float32
	FTPct          float32
	FG3Pct         float32
	Assists        int
	Reb            int
	Tov            int
}

func (s *Scoreboard) UnmarshalJSON(b []byte) error {
	var (
		sc      scoreboard
		headers []GameHeader
		lscores []LineScore
	)

	if err := json.Unmarshal(b, &sc); err != nil {
		return err
	}

	for i := range sc.ResultSets {
		var dest interface{}

		switch sc.ResultSets[i].Name {
		case "GameHeader":
			headers = make([]GameHeader, len(sc.ResultSets[i].RowSet))
			dest = headers
		case "LineScore":
			lscores = make([]LineScore, len(sc.ResultSets[i].RowSet))
			dest = lscores
		}

		if dest != nil {
			if err := decodeResultSet(dest, sc.ResultSets[i].RowSet); err != nil {
				return err
			}
		}
	}

	s.Date = sc.Parameters.GameDate
	s.Games = make([]Game, len(headers))

	for i := range s.Games {
		s.Games[i].Header = headers[i]
		s.Games[i].LineScore = lscores[i]
	}

	return nil
}

func decodeResultSet(dest interface{}, rowset []json.RawMessage) error {
	var (
		v = reflect.ValueOf(dest)
		t = v.Type().Elem()
	)

	for i := range rowset {
		sl := make([]interface{}, t.NumField())

		for j := 0; j < t.NumField(); j++ {
			sl[j] = v.Index(i).Field(j).Addr().Interface()
		}

		if err := json.Unmarshal(rowset[i], &sl); err != nil {
			return err
		}
	}

	return nil
}
