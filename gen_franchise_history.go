package baller

import (
	"net/url"
)

type FranchiseHistoryOptions struct {
	LeagueID string
}

type FranchiseHistoryResponse struct {
	FranchiseHistory []FranchiseHistoryFranchiseHistory
	DefunctTeams     []FranchiseHistoryDefunctTeams
}

func (c *Client) FranchiseHistory(options *FranchiseHistoryOptions) (*FranchiseHistoryResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "franchisehistory?"
		dest FranchiseHistoryResponse
		res  result
	)

	q.Set("LeagueID", encodeString(options.LeagueID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"FranchiseHistory": &dest.FranchiseHistory,
		"DefunctTeams":     &dest.DefunctTeams,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type FranchiseHistoryFranchiseHistory struct {
	LeagueID      string  `header:"LEAGUE_ID"`
	TeamID        int     `header:"TEAM_ID"`
	TeamCity      string  `header:"TEAM_CITY"`
	TeamName      string  `header:"TEAM_NAME"`
	StartYear     string  `header:"START_YEAR"`
	EndYear       string  `header:"END_YEAR"`
	Years         int     `header:"YEARS"`
	Games         int     `header:"GAMES"`
	Wins          int     `header:"WINS"`
	Losses        int     `header:"LOSSES"`
	WinPct        float32 `header:"WIN_PCT"`
	PoAppearances int     `header:"PO_APPEARANCES"`
	DivTitles     int     `header:"DIV_TITLES"`
	ConfTitles    int     `header:"CONF_TITLES"`
	LeagueTitles  int     `header:"LEAGUE_TITLES"`
}

type FranchiseHistoryDefunctTeams struct {
	LeagueID      string  `header:"LEAGUE_ID"`
	TeamID        int     `header:"TEAM_ID"`
	TeamCity      string  `header:"TEAM_CITY"`
	TeamName      string  `header:"TEAM_NAME"`
	StartYear     string  `header:"START_YEAR"`
	EndYear       string  `header:"END_YEAR"`
	Years         int     `header:"YEARS"`
	Games         int     `header:"GAMES"`
	Wins          int     `header:"WINS"`
	Losses        int     `header:"LOSSES"`
	WinPct        float32 `header:"WIN_PCT"`
	PoAppearances int     `header:"PO_APPEARANCES"`
	DivTitles     int     `header:"DIV_TITLES"`
	ConfTitles    int     `header:"CONF_TITLES"`
	LeagueTitles  int     `header:"LEAGUE_TITLES"`
}
