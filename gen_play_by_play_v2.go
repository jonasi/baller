package baller

import (
	"net/url"
)

type PlayByPlayV2Options struct {
	GameID      string
	StartPeriod int
	EndPeriod   int
}

type PlayByPlayV2Response struct {
	PlayByPlay     []PlayByPlayV2PlayByPlay
	AvailableVideo []PlayByPlayV2AvailableVideo
}

func (c *Client) PlayByPlayV2(options *PlayByPlayV2Options) (*PlayByPlayV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "playbyplayv2?"
		dest PlayByPlayV2Response
		res  result
	)

	q.Set("GameID", encodeString(options.GameID))
	q.Set("StartPeriod", encodeInt(options.StartPeriod))
	q.Set("EndPeriod", encodeInt(options.EndPeriod))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"PlayByPlay":     &dest.PlayByPlay,
		"AvailableVideo": &dest.AvailableVideo,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type PlayByPlayV2PlayByPlay struct {
	GameID                  string      `header:"GAME_ID"`
	Eventnum                int         `header:"EVENTNUM"`
	Eventmsgtype            int         `header:"EVENTMSGTYPE"`
	Eventmsgactiontype      int         `header:"EVENTMSGACTIONTYPE"`
	Period                  int         `header:"PERIOD"`
	Wctimestring            string      `header:"WCTIMESTRING"`
	Pctimestring            string      `header:"PCTIMESTRING"`
	Homedescription         string      `header:"HOMEDESCRIPTION"`
	Neutraldescription      interface{} `header:"NEUTRALDESCRIPTION"`
	Visitordescription      string      `header:"VISITORDESCRIPTION"`
	Score                   string      `header:"SCORE"`
	Scoremargin             string      `header:"SCOREMARGIN"`
	Person1type             int         `header:"PERSON1TYPE"`
	Player1ID               int         `header:"PLAYER1_ID"`
	Player1Name             string      `header:"PLAYER1_NAME"`
	Player1TeamID           int         `header:"PLAYER1_TEAM_ID"`
	Player1TeamCity         string      `header:"PLAYER1_TEAM_CITY"`
	Player1TeamNickname     string      `header:"PLAYER1_TEAM_NICKNAME"`
	Player1TeamAbbreviation string      `header:"PLAYER1_TEAM_ABBREVIATION"`
	Person2type             int         `header:"PERSON2TYPE"`
	Player2ID               int         `header:"PLAYER2_ID"`
	Player2Name             string      `header:"PLAYER2_NAME"`
	Player2TeamID           int         `header:"PLAYER2_TEAM_ID"`
	Player2TeamCity         string      `header:"PLAYER2_TEAM_CITY"`
	Player2TeamNickname     string      `header:"PLAYER2_TEAM_NICKNAME"`
	Player2TeamAbbreviation string      `header:"PLAYER2_TEAM_ABBREVIATION"`
	Person3type             int         `header:"PERSON3TYPE"`
	Player3ID               int         `header:"PLAYER3_ID"`
	Player3Name             string      `header:"PLAYER3_NAME"`
	Player3TeamID           int         `header:"PLAYER3_TEAM_ID"`
	Player3TeamCity         string      `header:"PLAYER3_TEAM_CITY"`
	Player3TeamNickname     string      `header:"PLAYER3_TEAM_NICKNAME"`
	Player3TeamAbbreviation string      `header:"PLAYER3_TEAM_ABBREVIATION"`
}

type PlayByPlayV2AvailableVideo struct {
	VideoAvailableFlag int `header:"VIDEO_AVAILABLE_FLAG"`
}
