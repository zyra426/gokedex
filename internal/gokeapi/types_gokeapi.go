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

type PokemonResp struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	BaseExperience int       `json:"base_experience"`
	Height         int       `json:"height"`
	IsDefault      bool      `json:"is_default"`
	Order          int       `json:"order"`
	Weight         int       `json:"weight"`
	Abilities      []Ability `json:"abilities"`
	Forms          []Result  `json:"forms"`
	GameIndices    []struct {
		GameIndex int    `json:"game_index"`
		Version   Result `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item           Result `json:"item"`
		VersionDetails []struct {
			Rarity  int    `json:"rarity"`
			Version Result `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                Result `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int    `json:"level_learned_at"`
			VersionGroup    Result `json:"version_group"`
			MoveLearnMethod Result `json:"move_learn_method"`
		} `json:"version_group_details"`
		Order int `json:"order"`
	} `json:"moves"`
	Species Result `json:"species"`
	Cries   struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int    `json:"base_stat"`
		Effort   int    `json:"effort"`
		Stat     Result `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int    `json:"slot"`
		Type Result `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation Result `json:"generation"`
		Types      []Type `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation Result    `json:"generation"`
		Abilities  []Ability `json:"abilities"`
	} `json:"past_abilities"`
}

type Type struct {
	Slot int    `json:"slot"`
	Type Result `json:"type"`
}

type Ability struct {
	IsHidden bool   `json:"is_hidden"`
	Slot     int    `json:"slot"`
	Ability  Result `json:"ability"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
