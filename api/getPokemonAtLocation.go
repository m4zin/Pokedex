package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type EncounterResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func GetPokemonsAtLocation(url string) (EncounterResponse, error) {
	cachedData, ok := SharedCache.Get(url)

	if ok {
		locationAreas := EncounterResponse{}
		if err := json.Unmarshal(cachedData, &locationAreas); err != nil {
			return EncounterResponse{}, fmt.Errorf("error unmarshaling json: %w", err)
		}
		return locationAreas, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return EncounterResponse{}, fmt.Errorf("error making get request: %w", err)
	}
	defer res.Body.Close()

	// NOTE: read body into a []byte for later unmarshaling into struct
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return EncounterResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	// NOTE: check if the request actually succeeded
	if res.StatusCode != http.StatusOK {
		return EncounterResponse{}, fmt.Errorf("unexpected status code: %d %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	SharedCache.Add(url, body)

	locationAreas := EncounterResponse{}
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return EncounterResponse{}, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return locationAreas, nil
}
