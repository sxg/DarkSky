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
	testURL    = "https://api.forecast.io/forecast/c00106496bfdbb8374474a24f7371e2d/41.4993,81.6944"
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
