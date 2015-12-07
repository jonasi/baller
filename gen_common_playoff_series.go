package baller

import (
	"net/url"
)

type CommonPlayoffSeriesOptions struct {
	LeagueID string
	Season   string
	SeriesID string
}

type CommonPlayoffSeriesResponse struct {
	PlayoffSeries []CommonPlayoffSeriesPlayoffSeries
}

func (c *Client) CommonPlayoffSeries(options *CommonPlayoffSeriesOptions) (*CommonPlayoffSeriesResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonplayoffseries?"
		dest CommonPlayoffSeriesResponse
		res  result
	)

	q.Set("LeagueID", encodeString(options.LeagueID))
	q.Set("Season", encodeString(options.Season))
	q.Set("SeriesID", encodeString(options.SeriesID))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"PlayoffSeries": &dest.PlayoffSeries,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type CommonPlayoffSeriesPlayoffSeries struct {
	GameID        string `header:"GAME_ID"`
	HomeTeamID    int    `header:"HOME_TEAM_ID"`
	VisitorTeamID int    `header:"VISITOR_TEAM_ID"`
	SeriesID      string `header:"SERIES_ID"`
	GameNum       int    `header:"GAME_NUM"`
}
