package baller

import (
	"net/url"
)

type CommonTeamYearsOptions struct {
	LeagueID string
}

type CommonTeamYearsResponse struct {
	TeamYears []CommonTeamYearsTeamYears
}

func (c *Client) CommonTeamYears(options *CommonTeamYearsOptions) (*CommonTeamYearsResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonteamyears?"
		dest CommonTeamYearsResponse
		res  result
	)

	q.Set("LeagueID", encodeString(options.LeagueID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"TeamYears": &dest.TeamYears,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type CommonTeamYearsTeamYears struct {
	LeagueID     string `header:"LEAGUE_ID"`
	TeamID       int    `header:"TEAM_ID"`
	MinYear      string `header:"MIN_YEAR"`
	MaxYear      string `header:"MAX_YEAR"`
	Abbreviation string `header:"ABBREVIATION"`
}
