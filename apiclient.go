package darksky

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// APIClient to interact with the Dark Sky API.
type APIClient struct {
	apiKey   string
	Units    Units
	Language Language
}

const (
	darkSkyURL = "https://api.forecast.io/forecast"
)

// NewAPIClient creates and returns a reference to a new APIClient to interact with the Dark Sky API.
func NewAPIClient(apiKey string) *APIClient {
	return &APIClient{apiKey: apiKey}
}

// GetForecast gets the current forecast at a specified latitude and longitude.
func (client *APIClient) GetForecast(lat, lon float64) (*Forecast, error) {
	// Create URL string from host, API key, and lat/lon
	url, err := buildURL(client.apiKey, lat, lon)
	if err != nil {
		return nil, err
	}

	// Make a GET request to the URL
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the body into a Forecast
	var forecast Forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return nil, err
	}
	return &forecast, nil
}

func buildURL(apiKey string, lat, lon float64) (*url.URL, error) {
	var latLonPath = strings.Join([]string{strconv.FormatFloat(lat, 'f', 6, 64), strconv.FormatFloat(lon, 'f', 6, 64)}, ",")
	var urlString = strings.Join([]string{darkSkyURL, apiKey, latLonPath}, "/")
	var url, err = url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	return url, nil
}
