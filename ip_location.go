package txlbs

import (
	"context"
	"net/url"
)

// IPLocationResponse IP定位回复
type IPLocationResponse struct {
	Meta
	Result struct {
		IP       string `json:"ip"`
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		AdInfo struct {
			Nation   string `json:"nation"`
			Province string `json:"province"`
			City     string `json:"city"`
			Adcode   int    `json:"adcode"`
		} `json:"ad_info"`
	} `json:"result"`
}

// IPLocation 定位查询
func (c *Client) IPLocation(ctx context.Context, ip string) (*IPLocationResponse, error) {
	v := url.Values{}
	if ip != "" {
		v.Set("ip", ip)
	}
	v, err := c.signatureQueryParameter(v, string(ipLocation))
	if err != nil {
		return nil, err
	}
	r := c.ca.NewRequest().WithPath(ipLocation.Full(v.Encode()))
	resp, err := c.ca.Do(ctx, r)
	if err != nil {
		return nil, err
	}
	var dr IPLocationResponse
	if err := resp.DecodeFromJSON(&dr); err != nil {
		return nil, err
	}
	return &dr, nil
}
