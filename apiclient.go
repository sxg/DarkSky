package DarkSky

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIClient to interact with the Dark Sky API.
type APIClient struct {
	Units    Units
	Language Language
}

const (
	darkSkyURL = "https://api.forecast.io/forecast"
)

func (client *APIClient) GetForecast(lat, lon float64) (*Forecast, error) {
	resp, err := http.Get(testURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var forecast Forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return nil, err
	}
	return &forecast, nil
}
