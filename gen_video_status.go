package baller

import (
	"net/url"
)

type VideoStatusOptions struct {
	LeagueID string
	GameDate string
}

type VideoStatusResponse struct {
	VideoStatus []VideoStatusVideoStatus
}

func (c *Client) VideoStatus(options *VideoStatusOptions) (*VideoStatusResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "videostatus?"
		dest VideoStatusResponse
		res  result
	)

	q.Set("LeagueID", encodeString(options.LeagueID))
	q.Set("GameDate", encodeString(options.GameDate))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"VideoStatus": &dest.VideoStatus,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type VideoStatusVideoStatus struct {
	GameID                  string `header:"GAME_ID"`
	GameDate                string `header:"GAME_DATE"`
	VisitorTeamID           int    `header:"VISITOR_TEAM_ID"`
	VisitorTeamCity         string `header:"VISITOR_TEAM_CITY"`
	VisitorTeamName         string `header:"VISITOR_TEAM_NAME"`
	VisitorTeamAbbreviation string `header:"VISITOR_TEAM_ABBREVIATION"`
	HomeTeamID              int    `header:"HOME_TEAM_ID"`
	HomeTeamCity            string `header:"HOME_TEAM_CITY"`
	HomeTeamName            string `header:"HOME_TEAM_NAME"`
	HomeTeamAbbreviation    string `header:"HOME_TEAM_ABBREVIATION"`
	GameStatus              int    `header:"GAME_STATUS"`
	GameStatusText          string `header:"GAME_STATUS_TEXT"`
	IsAvailable             int    `header:"IS_AVAILABLE"`
	PtXyzAvailable          int    `header:"PT_XYZ_AVAILABLE"`
}
