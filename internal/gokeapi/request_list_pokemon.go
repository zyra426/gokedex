package gokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) FindPokemon(name string) (PokemonResp, error) {
	url := baseURL + "/pokemon/" + name

	if ce, exists := c.cache.Get(url); exists {
		pokemonResp := PokemonResp{}
		err := json.Unmarshal(ce, &pokemonResp)
		if err != nil {
			return pokemonResp, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemonResp := PokemonResp{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return pokemonResp, err
	}

	return pokemonResp, nil
}
