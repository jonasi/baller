package baller

import (
	"net/url"
)

type BoxscoreTraditionalV2Options struct {
}

type BoxscoreTraditionalV2Response struct {
}

func (c *Client) BoxscoreTraditionalV2(options *BoxscoreTraditionalV2Options) (*BoxscoreTraditionalV2Response, error) {
	var (
		q    = url.Values{}
		url  = baseURL + "boxscoretranditionalv2?"
		dest BoxscoreTraditionalV2Response
		res  result
	)

	if err := c.do(url+q.Encode(), &res); err != nil {
		return nil, err
	}

	err := res.unmarshalResultSets(map[string]interface{}{})

	if err != nil {
		return nil, err
	}

	return &dest, nil
}
