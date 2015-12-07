package baller

import (
	"net/url"
)

type PlayByPlayOptions struct {
	GameID      string
	StartPeriod int
	EndPeriod   int
}

type PlayByPlayResponse struct {
	PlayByPlay     []PlayByPlayPlayByPlay
	AvailableVideo []PlayByPlayAvailableVideo
}

func (c *Client) PlayByPlay(options *PlayByPlayOptions) (*PlayByPlayResponse, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "playbyplay?"
		dest PlayByPlayResponse
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

type PlayByPlayPlayByPlay struct {
	GameID             string      `header:"GAME_ID"`
	Eventnum           int         `header:"EVENTNUM"`
	Eventmsgtype       int         `header:"EVENTMSGTYPE"`
	Eventmsgactiontype int         `header:"EVENTMSGACTIONTYPE"`
	Period             int         `header:"PERIOD"`
	Wctimestring       string      `header:"WCTIMESTRING"`
	Pctimestring       string      `header:"PCTIMESTRING"`
	Homedescription    string      `header:"HOMEDESCRIPTION"`
	Neutraldescription interface{} `header:"NEUTRALDESCRIPTION"`
	Visitordescription string      `header:"VISITORDESCRIPTION"`
	Score              string      `header:"SCORE"`
	Scoremargin        string      `header:"SCOREMARGIN"`
}

type PlayByPlayAvailableVideo struct {
	VideoAvailableFlag int `header:"VIDEO_AVAILABLE_FLAG"`
}
