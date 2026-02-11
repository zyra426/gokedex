package gokeapi

type LocationResp struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type AreaExploreResp struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Result `json:"pokemon"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
