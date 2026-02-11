package gokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationResp struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string
	URL  string
}

func (c *Client) ListLocations(pageURL *string) (LocationResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if ce, exists := c.cache.Get(url); exists {
		locationsResp := LocationResp{}
		if err := json.Unmarshal(ce, &locationsResp); err != nil {
			return locationsResp, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResp{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResp{}, err
	}

	locationsResp := LocationResp{}
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return locationsResp, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}
