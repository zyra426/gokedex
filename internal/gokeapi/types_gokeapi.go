package gokeapi

type LocationResp struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type AreaExploreResp struct {
	EncounterMethodRates []struct {
		EncounterMethod Result `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int    `json:"rate"`
			Version Result `json:"version_details"`
		}
	} `json:"encounter_method_rates"`
	GameIndex int    `json:"game_index"`
	ID        int    `json:"id"`
	Location  Result `json:"location"`
	Name      string `json:"name"`
	Names     []struct {
		Language Result `json:"language"`
		Name     string `json:"name"`
	} `json:"names"`
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon        Result `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int           `json:"chance"`
			ConditionValues []interface{} `json:"condition_values"`
			MaxLevel        int           `json:"max_level"`
			Method          Result        `json:"method"`
			MinLevel        int           `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int    `json:"max_chance"`
		Version   Result `json:"version"`
	} `json:"version_details"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
