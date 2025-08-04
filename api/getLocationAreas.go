package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreasResponse struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous *string         `json:"previous"`
	Results  []LocationAreas `json:"results"`
}

func GetLocationAreas(url string) (LocationAreasResponse, error) {
	cachedData, ok := SharedCache.Get(url)

	if ok {
		locationAreas := LocationAreasResponse{}
		if err := json.Unmarshal(cachedData, &locationAreas); err != nil {
			return LocationAreasResponse{}, fmt.Errorf("error unmarshaling json: %w", err)
		}
		return locationAreas, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error making get request: %w", err)
	}
	defer res.Body.Close()

	// NOTE: read body into a []byte for later unmarshaling into struct
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("unexpected status code: %d %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	SharedCache.Add(url, body)

	locationAreas := LocationAreasResponse{}
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return locationAreas, nil
}
