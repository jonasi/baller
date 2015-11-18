package baller

import (
	"encoding/json"
	"fmt"
	"reflect"
)

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

type GameV2 struct {
	Game
	HomeTeamLeaders     TeamLeaders
	VisitingTeamLeaders TeamLeaders
	TicketLinks         TicketLinks
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

type TeamLeaders struct {
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

type TicketLinks struct {
	GameID    string `header:"GAME_ID"`
	LeagueTix string `header:"LEAG_TIX"`
}

type Scoreboard struct {
	Date  string
	Games []Game
}

func (s *Scoreboard) UnmarshalJSON(b []byte) error {
	var res result

	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}

	g, err := unmarshalScoreboardV1(&res)

	if err != nil {
		return err
	}

	s.Date = res.Parameters["GameDate"].(string)
	s.Games = g

	return nil
}

type ScoreboardV2 struct {
	Date  string
	Games []GameV2
}

func (s *ScoreboardV2) UnmarshalJSON(b []byte) error {
	var res result

	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}

	g, err := unmarshalScoreboardV1(&res)

	if err != nil {
		return err
	}

	s.Date = res.Parameters["GameDate"].(string)
	s.Games = make([]GameV2, len(g))

	for i := range g {
		s.Games[i].Game = g[i]
	}

	var (
		tickets []TicketLinks
		leaders []TeamLeaders
	)

	for i := range res.ResultSets {
		var (
			dest interface{}
			rs   = res.ResultSets[i]
			l    = len(rs.RowSet)
		)

		switch res.ResultSets[i].Name {
		case "TeamLeaders":
			leaders = make([]TeamLeaders, l)
			dest = leaders
		case "TicketLinks":
			tickets = make([]TicketLinks, l)
			dest = tickets
		}

		if dest != nil {
			if err := decodeResultSet(dest, rs.Headers, rs.RowSet); err != nil {
				return err
			}
		}
	}

	gameMap := map[string]*GameV2{}

	for i := range s.Games {
		gameMap[s.Games[i].Header.GameID] = &s.Games[i]
	}

	for i := range tickets {
		gm := gameMap[tickets[i].GameID]
		gm.TicketLinks = tickets[i]
	}

	for i := range leaders {
		gm := gameMap[leaders[i].GameID]

		if leaders[i].TeamID == gm.Header.HomeTeamID {
			gm.HomeTeamLeaders = leaders[i]
		} else {
			gm.VisitingTeamLeaders = leaders[i]
		}
	}

	return nil
}
func unmarshalScoreboardV1(res *result) ([]Game, error) {
	var (
		headers   []GameHeader
		lscores   []LineScore
		standings []SeriesStandings
		last      []LastMeeting
		est       []Standings
		wst       []Standings
		av        []Available
	)

	for i := range res.ResultSets {
		var (
			dest interface{}
			rs   = res.ResultSets[i]
			l    = len(rs.RowSet)
		)

		switch res.ResultSets[i].Name {
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
				return nil, err
			}
		}
	}

	games := make([]Game, len(headers))
	gameMap := map[string]*Game{}

	for i := range games {
		gameMap[headers[i].GameID] = &games[i]
		games[i].Header = headers[i]
		games[i].EasternConferenceStandings = est
		games[i].WesternConferenceStandings = wst
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

	return games, nil
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
