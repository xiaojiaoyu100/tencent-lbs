package txlbs

import (
	"context"
)

// DistrictResponse 行政区划回复
type DistrictResponse struct {
	Status      int               `json:"status"`
	Message     string            `json:"message"`
	DataVersion string            `json:"data_version"`
	Result      [][]*DistrictInfo `json:"result"`
}

// DistrictInfo 行政区划信息
type DistrictInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullname"`
	Location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
	Pinyin []string `json:"pinyin"`
	Cidx   []int    `json:"Cidx"`
}

// GetDistrictList 获取行政区划信息
func (c *Client) GetDistrictList(ctx context.Context) (*DistrictResponse, error) {
	var param struct {
		Key string `url:"key"`
	}
	param.Key = c.Config.Key
	r := c.ca.NewRequest().WithPath(districtList).WithQueryParam(&param)
	resp, err := c.ca.Do(ctx, r)
	if err != nil {
		return nil, err
	}
	var dr DistrictResponse
	if err := resp.DecodeFromJSON(&dr); err != nil {
		return nil, err
	}
	return &dr, nil
}

// GetDistrictChildren 根据子级行政区划
func (c *Client) GetDistrictChildren(ctx context.Context, id string) (*DistrictResponse, error) {
	var param struct {
		Key string `url:"key"`
		// 父级行政区划ID，缺省时则返回最顶级行政区划
		ID string `url:"id,omitempty"`
	}
	if id != "" {
		param.ID = id
	}
	param.Key = c.Config.Key
	r := c.ca.NewRequest().WithPath(districtGetChildren).WithQueryParam(&param)
	resp, err := c.ca.Do(ctx, r)
	if err != nil {
		return nil, err
	}
	var dr DistrictResponse
	if err := resp.DecodeFromJSON(&dr); err != nil {
		return nil, err
	}
	return &dr, nil
}

// DistrictSearch 行政区划搜索
func (c *Client) DistrictSearch(ctx context.Context, keyword string) (*DistrictResponse, error) {
	var param struct {
		Key string `url:"key"`
		// 搜索关键词：
		// 1.支持输入一个文本关键词
		// 2.支持多个行政区划代码，英文逗号分隔
		Keyword string `url:"keyword"`
	}
	if keyword != "" {
		param.Keyword = keyword
	}
	param.Key = c.Config.Key
	r := c.ca.NewRequest().WithPath(districtSearch).WithQueryParam(&param)
	resp, err := c.ca.Do(ctx, r)
	if err != nil {
		return nil, err
	}
	var dr DistrictResponse
	if err := resp.DecodeFromJSON(&dr); err != nil {
		return nil, err
	}
	return &dr, nil
}
