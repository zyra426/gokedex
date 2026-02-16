package gokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/buntdb"
)

func (c *Client) CatchPokemon(name string) (PokemonResp, error) {
	url := baseURL + "/pokemon/" + name

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
		return PokemonResp{}, err
	}

	setKey := fmt.Sprintf("pokemon:%s", pokemonResp.Name)
	setVal, err := json.Marshal(pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	dbErr := c.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(setKey, string(setVal), nil)
		return err
	})
	if dbErr != nil {
		return PokemonResp{}, err
	}

	return pokemonResp, nil
}
