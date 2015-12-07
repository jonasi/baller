package baller

import (
	"net/url"
)

type BoxscoreScoringV2Options struct {
	GameID      string
	StartPeriod int
	EndPeriod   int
	StartRange  int
	EndRange    int
	RangeType   int
}

type BoxscoreScoringV2Response struct {
	SqlPlayersScoring []BoxscoreScoringV2SqlPlayersScoring
	SqlTeamsScoring   []BoxscoreScoringV2SqlTeamsScoring
}

func (c *Client) BoxscoreScoringV2(options *BoxscoreScoringV2Options) (*BoxscoreScoringV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscorescoringv2?"
		dest BoxscoreScoringV2Response
		res  result
	)

	q.Set("GameID", encodeString(options.GameID))
	q.Set("StartPeriod", encodeInt(options.StartPeriod))
	q.Set("EndPeriod", encodeInt(options.EndPeriod))
	q.Set("StartRange", encodeInt(options.StartRange))
	q.Set("EndRange", encodeInt(options.EndRange))
	q.Set("RangeType", encodeInt(options.RangeType))

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{
		"sqlPlayersScoring": &dest.SqlPlayersScoring,
		"sqlTeamsScoring":   &dest.SqlTeamsScoring,
	})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BoxscoreScoringV2SqlPlayersScoring struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	PlayerID         int     `header:"PLAYER_ID"`
	PlayerName       string  `header:"PLAYER_NAME"`
	StartPosition    string  `header:"START_POSITION"`
	Comment          string  `header:"COMMENT"`
	Min              string  `header:"MIN"`
	PctFga2pt        float32 `header:"PCT_FGA_2PT"`
	PctFga3pt        float32 `header:"PCT_FGA_3PT"`
	PctPts2pt        float32 `header:"PCT_PTS_2PT"`
	PctPts2ptMr      float32 `header:"PCT_PTS_2PT_MR"`
	PctPts3pt        float32 `header:"PCT_PTS_3PT"`
	PctPtsFb         float32 `header:"PCT_PTS_FB"`
	PctPtsFt         float32 `header:"PCT_PTS_FT"`
	PctPtsOffTov     float32 `header:"PCT_PTS_OFF_TOV"`
	PctPtsPaint      float32 `header:"PCT_PTS_PAINT"`
	PctAst2pm        float32 `header:"PCT_AST_2PM"`
	PctUast2pm       float32 `header:"PCT_UAST_2PM"`
	PctAst3pm        float32 `header:"PCT_AST_3PM"`
	PctUast3pm       float32 `header:"PCT_UAST_3PM"`
	PctAstFgm        float32 `header:"PCT_AST_FGM"`
	PctUastFgm       float32 `header:"PCT_UAST_FGM"`
}

type BoxscoreScoringV2SqlTeamsScoring struct {
	GameID           string  `header:"GAME_ID"`
	TeamID           int     `header:"TEAM_ID"`
	TeamName         string  `header:"TEAM_NAME"`
	TeamAbbreviation string  `header:"TEAM_ABBREVIATION"`
	TeamCity         string  `header:"TEAM_CITY"`
	Min              string  `header:"MIN"`
	PctFga2pt        float32 `header:"PCT_FGA_2PT"`
	PctFga3pt        float32 `header:"PCT_FGA_3PT"`
	PctPts2pt        float32 `header:"PCT_PTS_2PT"`
	PctPts2ptMr      float32 `header:"PCT_PTS_2PT_MR"`
	PctPts3pt        float32 `header:"PCT_PTS_3PT"`
	PctPtsFb         float32 `header:"PCT_PTS_FB"`
	PctPtsFt         float32 `header:"PCT_PTS_FT"`
	PctPtsOffTov     float32 `header:"PCT_PTS_OFF_TOV"`
	PctPtsPaint      float32 `header:"PCT_PTS_PAINT"`
	PctAst2pm        float32 `header:"PCT_AST_2PM"`
	PctUast2pm       float32 `header:"PCT_UAST_2PM"`
	PctAst3pm        float32 `header:"PCT_AST_3PM"`
	PctUast3pm       float32 `header:"PCT_UAST_3PM"`
	PctAstFgm        float32 `header:"PCT_AST_FGM"`
	PctUastFgm       float32 `header:"PCT_UAST_FGM"`
}
