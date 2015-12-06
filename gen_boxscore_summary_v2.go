package baller

import (
	"net/url"
)

type BoxscoreSummaryV2Options struct {
	GameID string
}

type BoxscoreSummaryV2Response struct {
	GameSummary     []BoxscoreSummaryV2GameSummary
	OtherStats      []BoxscoreSummaryV2OtherStats
	Officials       []BoxscoreSummaryV2Officials
	InactivePlayers []BoxscoreSummaryV2InactivePlayers
	GameInfo        []BoxscoreSummaryV2GameInfo
	LineScore       []BoxscoreSummaryV2LineScore
	LastMeeting     []BoxscoreSummaryV2LastMeeting
	SeasonSeries    []BoxscoreSummaryV2SeasonSeries
	AvailableVideo  []BoxscoreSummaryV2AvailableVideo
}

func (c *Client) BoxscoreSummaryV2(options *BoxscoreSummaryV2Options) (*BoxscoreSummaryV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscoresummaryv2?"
		dest BoxscoreSummaryV2Response
		res  result
	)

	q.Set("GameID", encodeString(options.GameID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"GameSummary":     &dest.GameSummary,
		"OtherStats":      &dest.OtherStats,
		"Officials":       &dest.Officials,
		"InactivePlayers": &dest.InactivePlayers,
		"GameInfo":        &dest.GameInfo,
		"LineScore":       &dest.LineScore,
		"LastMeeting":     &dest.LastMeeting,
		"SeasonSeries":    &dest.SeasonSeries,
		"AvailableVideo":  &dest.AvailableVideo,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BoxscoreSummaryV2GameSummary struct {
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

type BoxscoreSummaryV2OtherStats struct {
	LeagueID         string `header:"LEAGUE_ID"`
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

type BoxscoreSummaryV2Officials struct {
	OfficialID int    `header:"OFFICIAL_ID"`
	FirstName  string `header:"FIRST_NAME"`
	LastName   string `header:"LAST_NAME"`
	JerseyNum  string `header:"JERSEY_NUM"`
}

type BoxscoreSummaryV2InactivePlayers struct {
	PlayerID         int    `header:"PLAYER_ID"`
	FirstName        string `header:"FIRST_NAME"`
	LastName         string `header:"LAST_NAME"`
	JerseyNum        string `header:"JERSEY_NUM"`
	TeamID           int    `header:"TEAM_ID"`
	TeamCity         string `header:"TEAM_CITY"`
	TeamName         string `header:"TEAM_NAME"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
}

type BoxscoreSummaryV2GameInfo struct {
	GameDate   string `header:"GAME_DATE"`
	Attendance int    `header:"ATTENDANCE"`
	GameTime   string `header:"GAME_TIME"`
}

type BoxscoreSummaryV2LineScore struct {
	GameDateEst      string `header:"GAME_DATE_EST"`
	GameSequence     int    `header:"GAME_SEQUENCE"`
	GameID           string `header:"GAME_ID"`
	TeamID           int    `header:"TEAM_ID"`
	TeamAbbreviation string `header:"TEAM_ABBREVIATION"`
	TeamCityName     string `header:"TEAM_CITY_NAME"`
	TeamNickname     string `header:"TEAM_NICKNAME"`
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

type BoxscoreSummaryV2LastMeeting struct {
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

type BoxscoreSummaryV2SeasonSeries struct {
	GameID         string `header:"GAME_ID"`
	HomeTeamID     int    `header:"HOME_TEAM_ID"`
	VisitorTeamID  int    `header:"VISITOR_TEAM_ID"`
	GameDateEst    string `header:"GAME_DATE_EST"`
	HomeTeamWins   int    `header:"HOME_TEAM_WINS"`
	HomeTeamLosses int    `header:"HOME_TEAM_LOSSES"`
	SeriesLeader   string `header:"SERIES_LEADER"`
}

type BoxscoreSummaryV2AvailableVideo struct {
	GameID             string `header:"GAME_ID"`
	VideoAvailableFlag int    `header:"VIDEO_AVAILABLE_FLAG"`
	PtAvailable        int    `header:"PT_AVAILABLE"`
	PtXyzAvailable     int    `header:"PT_XYZ_AVAILABLE"`
	WhStatus           int    `header:"WH_STATUS"`
}
