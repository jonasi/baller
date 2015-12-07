package baller

import (
	"net/url"
)

type CommonTeamRosterOptions struct {
	TeamID   int
	LeagueID string
	Season   string
}

type CommonTeamRosterResponse struct {
	CommonTeamRoster []CommonTeamRosterCommonTeamRoster
	Coaches          []CommonTeamRosterCoaches
}

func (c *Client) CommonTeamRoster(options *CommonTeamRosterOptions) (*CommonTeamRosterResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "commonteamroster?"
		dest CommonTeamRosterResponse
		res  result
	)

	q.Set("TeamID", encodeInt(options.TeamID))
	q.Set("LeagueID", encodeString(options.LeagueID))
	q.Set("Season", encodeString(options.Season))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"CommonTeamRoster": &dest.CommonTeamRoster,
		"Coaches":          &dest.Coaches,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type CommonTeamRosterCommonTeamRoster struct {
	Teamid    int     `header:"TeamID"`
	Season    string  `header:"SEASON"`
	Leagueid  string  `header:"LeagueID"`
	Player    string  `header:"PLAYER"`
	Num       string  `header:"NUM"`
	Position  string  `header:"POSITION"`
	Height    string  `header:"HEIGHT"`
	Weight    string  `header:"WEIGHT"`
	BirthDate string  `header:"BIRTH_DATE"`
	Age       float32 `header:"AGE"`
	Exp       string  `header:"EXP"`
	School    string  `header:"SCHOOL"`
	PlayerID  int     `header:"PLAYER_ID"`
}

type CommonTeamRosterCoaches struct {
	TeamID       int     `header:"TEAM_ID"`
	Season       string  `header:"SEASON"`
	CoachID      string  `header:"COACH_ID"`
	FirstName    string  `header:"FIRST_NAME"`
	LastName     string  `header:"LAST_NAME"`
	CoachName    string  `header:"COACH_NAME"`
	CoachCode    string  `header:"COACH_CODE"`
	IsAssistant  float32 `header:"IS_ASSISTANT"`
	CoachType    string  `header:"COACH_TYPE"`
	School       string  `header:"SCHOOL"`
	SortSequence float32 `header:"SORT_SEQUENCE"`
}
