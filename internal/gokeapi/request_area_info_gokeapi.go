package gokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListAreaInfo(areaName string) (AreaExploreResp, error) {
	url := baseURL + "/location-area/" + areaName

	if ce, exists := c.cache.Get(url); exists {
		areaResp := AreaExploreResp{}
		err := json.Unmarshal(ce, &areaResp)
		if err != nil {
			return AreaExploreResp{}, err
		}

		return areaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaExploreResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return AreaExploreResp{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaExploreResp{}, err
	}

	areaResp := AreaExploreResp{}
	if err := json.Unmarshal(data, &areaResp); err != nil {
		return areaResp, err
	}

	c.cache.Add(url, data)
	return areaResp, nil
}
