package baller

import (
	"encoding/json"
	"fmt"
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
	Available                  Available
	Header                     GameHeader
	HomeLineScore              LineScore
	VisitorLineScore           LineScore
	SeriesStandings            SeriesStandings
	LastMeeting                LastMeeting
	EasternConferenceStandings []Standings
	WesternConferenceStandings []Standings
}

type GameHeader struct {
	GameDateEST             string `header:"GAME_DATE_EST"`
	GameSequence            int    `header:"GAME_SEQUENCE"`
	GameID                  string `header:"GAME_ID"`
	GameStatusID            int    `header:"GAME_STATUS_ID"`
	GameStatusText          string `header:"GAME_STATUS_TEXT"`
	GameCode                string `header:"GAMECODE"`
	HomeTeamID              int    `header:"HOME_TEAM_ID"`
	VisitorTeamID           int    `header:"VISITOR_TEAM_ID"`
	Season                  string `header:"SEASON"`
	LivePeriod              int    `header:"LIVE_PERIOD"`
	LivePCTime              string `header:"LIVE_PC_TIME"`
	NatlTVBroadcasterAbbrev string `header:"NATL_TV_BROADCASTER_ABBREVIATION"`
	LivePeriodTimeBcast     string `header:"LIVE_PERIOD_TIME_BCAST"`
	WHStatus                int    `header:"WH_STATUS"`
}

type LineScore struct {
	GameDateEST    string  `header:"GAME_DATE_EST"`
	GameSequence   int     `header:"GAME_SEQUENCE"`
	GameID         string  `header:"GAME_ID"`
	TeamID         int     `header:"TEAM_ID"`
	TeamAbbrev     string  `header:"TEAM_ABBREVIATION"`
	TeamCityName   string  `header:"TEAM_CITY_NAME"`
	TeamWinsLosses string  `header:"TEAM_WINS_LOSSES"`
	PtsQtr1        int     `header:"PTS_QTR1"`
	PtsQtr2        int     `header:"PTS_QTR2"`
	PtsQtr3        int     `header:"PTS_QTR3"`
	PtsQtr4        int     `header:"PTS_QTR4"`
	PtsOT1         int     `header:"PTS_OT1"`
	PtsOT2         int     `header:"PTS_OT2"`
	PtsOT3         int     `header:"PTS_OT3"`
	PtsOT4         int     `header:"PTS_OT4"`
	PtsOT5         int     `header:"PTS_OT5"`
	PtsOT6         int     `header:"PTS_OT6"`
	PtsOT7         int     `header:"PTS_OT7"`
	PtsOT8         int     `header:"PTS_OT8"`
	PtsOT9         int     `header:"PTS_OT9"`
	PtsOT10        int     `header:"PTS_OT10"`
	Pts            int     `header:"PTS"`
	FGPct          float32 `header:"FG_PCT"`
	FTPct          float32 `header:"FT_PCT"`
	FG3Pct         float32 `header:"FG3_PCT"`
	Assists        int     `header:"AST"`
	Reb            int     `header:"REB"`
	Tov            int     `header:"TOV"`
}

type SeriesStandings struct {
	GameID         string `header:"GAME_ID"`
	HomeTeamID     int    `header:"HOME_TEAM_ID"`
	VisitorTeamID  int    `header:"VISITOR_TEAM_ID"`
	GameDateEST    string `header:"GAME_DATE_EST"`
	HomeTeamWins   int    `header:"HOME_TEAM_WINS"`
	HomeTeamLosses int    `header:"HOME_TEAM_LOSSES"`
	SeriesLeader   string `header:"SERIES_LEADER"`
}

type LastMeeting struct {
	GameID                          string `header:"GAME_ID"`
	LastGameID                      string `header:"LAST_GAME_ID"`
	LastGameDateEST                 string `header:"LAST_GAME_DATE_EST"`
	LastGameHomeTeamID              int    `header:"LAST_GAME_HOME_TEAM_ID"`
	LastGameHomeTeamCity            string `header:"LAST_GAME_HOME_TEAM_CITY"`
	LastGameHomeTeamName            string `header:"LAST_GAME_HOME_TEAM_NAME"`
	LastGameHomeTeamAbbreviation    string `header:"LAST_GAME_HOME_TEAM_ABBREVIATION"`
	LastGameHomeTeamPoints          int    `header:"LAST_GAME_HOME_TEAM_POINTS"`
	LastGameVisitorTeamID           int    `header:"LAST_GAME_VISITOR_TEAM_ID"`
	LastGameVisitorTeamCity         string `header:"LAST_GAME_VISITOR_TEAM_CITY"`
	LastGameVisitorTeamName         string `header:"LAST_GAME_VISITOR_TEAM_NAME"`
	LastGameVisitorTeamAbbreviation string `header:"LAST_GAME_VISITOR_TEAM_CITY1"`
	LastGameVisitorTeamPoints       int    `header:"LAST_GAME_VISITOR_TEAM_POINTS"`
}

type Standings struct {
	TeamID        int     `header:"TEAM_ID"`
	LeagueID      string  `header:"LEAGUE_ID"`
	SeasonID      string  `header:"SEASON_ID"`
	StandingsDate string  `header:"STANDINGSDATE"`
	Conference    string  `header:"CONFERENCE"`
	Team          string  `header:"TEAM"`
	Game          int     `header:"GAME"`
	Wins          int     `header:"W"`
	Losses        int     `header:"L"`
	Pct           float32 `header:"W_PCT"`
	HomeRecord    string  `header:"HOME_RECORD"`
	RoadRecord    string  `header:"ROAD_RECORD"`
}

type Available struct {
	GameID      string `header:"GAME_ID"`
	PTAvailable int    `header:"PT_AVAILABLE"`
}

func (s *Scoreboard) UnmarshalJSON(b []byte) error {
	var (
		sc        scoreboard
		headers   []GameHeader
		lscores   []LineScore
		standings []SeriesStandings
		last      []LastMeeting
		est       []Standings
		wst       []Standings
		av        []Available
	)

	if err := json.Unmarshal(b, &sc); err != nil {
		return err
	}

	for i := range sc.ResultSets {
		var (
			dest interface{}
			rs   = sc.ResultSets[i]
			l    = len(rs.RowSet)
		)

		switch sc.ResultSets[i].Name {
		case "GameHeader":
			headers = make([]GameHeader, l)
			dest = headers
		case "LineScore":
			lscores = make([]LineScore, l)
			dest = lscores
		case "SeriesStandings":
			standings = make([]SeriesStandings, l)
			dest = standings
		case "LastMeeting":
			last = make([]LastMeeting, l)
			dest = last
		case "EastConfStandingsByDay":
			est = make([]Standings, l)
			dest = est
		case "WestConfStandingsByDay":
			wst = make([]Standings, l)
			dest = wst
		case "Available":
			av = make([]Available, l)
			dest = av
		}

		if dest != nil {
			if err := decodeResultSet(dest, rs.Headers, rs.RowSet); err != nil {
				return err
			}
		}
	}

	s.Date = sc.Parameters.GameDate
	s.Games = make([]Game, len(headers))
	gameMap := map[string]*Game{}

	for i := range s.Games {
		gameMap[headers[i].GameID] = &s.Games[i]
		s.Games[i].Header = headers[i]
		s.Games[i].EasternConferenceStandings = est
		s.Games[i].WesternConferenceStandings = wst
	}

	for i := range av {
		gm := gameMap[av[i].GameID]
		gm.Available = av[i]
	}

	for i := range standings {
		gm := gameMap[standings[i].GameID]
		gm.SeriesStandings = standings[i]
	}

	for i := range last {
		gm := gameMap[last[i].GameID]
		gm.LastMeeting = last[i]
	}

	for i := range lscores {
		gm := gameMap[lscores[i].GameID]

		if lscores[i].TeamID == gm.Header.HomeTeamID {
			gm.HomeLineScore = lscores[i]
		} else {
			gm.VisitorLineScore = lscores[i]
		}
	}

	return nil
}

func decodeResultSet(dest interface{}, headers []string, rowset []json.RawMessage) error {
	var (
		v  = reflect.ValueOf(dest)
		t  = v.Type().Elem()
		mp = mkHeaderMap(t)
	)

	if len(headers) != len(mp) {
		return fmt.Errorf("Expected %d headers, found %d for type %T\nheaders: %#v\nmap: %#v", len(headers), len(mp), dest, headers, mp)
	}

	for i := range rowset {
		sl := make([]interface{}, len(headers))

		for j, h := range headers {
			if idx, ok := mp[h]; ok {
				sl[j] = v.Index(i).Field(idx).Addr().Interface()
			}
		}

		if err := json.Unmarshal(rowset[i], &sl); err != nil {
			return err
		}
	}

	return nil
}

func mkHeaderMap(t reflect.Type) map[string]int {
	mp := map[string]int{}

	for i := 0; i < t.NumField(); i++ {
		h := t.Field(i).Tag.Get("header")

		if h == "" {
			continue
		}

		mp[h] = i
	}

	return mp
}
