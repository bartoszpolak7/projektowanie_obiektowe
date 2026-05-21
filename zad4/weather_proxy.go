package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherProxy struct{ URL string }

func (p *WeatherProxy) FetchStation(station string) (Weather, error) {
	requestURL := p.URL + station
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return Weather{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Weather{}, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			return Weather{}, fmt.Errorf("station not found: %s", station)
		}
		return Weather{}, fmt.Errorf("failed to fetch data for station: %s", station)
	}

	var weather Weather

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		resp.Body.Close()
		return Weather{}, fmt.Errorf("failed to decode response for station: %s: %w", station, err)
	}

	resp.Body.Close()
	return weather, nil
}
