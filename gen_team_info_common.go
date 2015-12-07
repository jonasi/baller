package baller

import (
	"net/url"
)

type TeamInfoCommonOptions struct {
	LeagueID   string
	Season     string
	SeasonType string
	TeamID     int
}

type TeamInfoCommonResponse struct {
	TeamInfoCommon  []TeamInfoCommonTeamInfoCommon
	TeamSeasonRanks []TeamInfoCommonTeamSeasonRanks
}

func (c *Client) TeamInfoCommon(options *TeamInfoCommonOptions) (*TeamInfoCommonResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "teaminfocommon?"
		dest TeamInfoCommonResponse
		res  result
	)

	q.Set("LeagueID", encodeString(options.LeagueID))
	q.Set("Season", encodeString(options.Season))
	q.Set("SeasonType", encodeString(options.SeasonType))
	q.Set("TeamID", encodeInt(options.TeamID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"TeamInfoCommon":  &dest.TeamInfoCommon,
		"TeamSeasonRanks": &dest.TeamSeasonRanks,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type TeamInfoCommonTeamInfoCommon struct {
	TeamID           int     `header:"TEAM_ID"`
	SeasonYear       string  `header:"SEASON_YEAR"`
	TeamCity         string  `header:"TEAM_CITY"`
	TeamName         string  `header:"TEAM_NAME"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamConference   string  `header:"TEAM_CONFERENCE"`
	TeamDivision     string  `header:"TEAM_DIVISION"`
	TeamCode         string  `header:"TEAM_CODE"`
	W                int     `header:"W"`
	L                int     `header:"L"`
	Pct              float32 `header:"PCT"`
	ConfRank         int     `header:"CONF_RANK"`
	DivRank          int     `header:"DIV_RANK"`
	MinYear          string  `header:"MIN_YEAR"`
	MaxYear          string  `header:"MAX_YEAR"`
}

type TeamInfoCommonTeamSeasonRanks struct {
	LeagueID   string  `header:"LEAGUE_ID"`
	SeasonID   string  `header:"SEASON_ID"`
	TeamID     int     `header:"TEAM_ID"`
	PtsRank    int     `header:"PTS_RANK"`
	PtsPg      float32 `header:"PTS_PG"`
	RebRank    int     `header:"REB_RANK"`
	RebPg      float32 `header:"REB_PG"`
	AstRank    int     `header:"AST_RANK"`
	AstPg      float32 `header:"AST_PG"`
	OppPtsRank int     `header:"OPP_PTS_RANK"`
	OppPtsPg   float32 `header:"OPP_PTS_PG"`
}
