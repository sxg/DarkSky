package DarkSky_test

import (
	"os"
	"testing"

	"github.com/sxg/DarkSky"
)

func TestNewAPIClient(t *testing.T) {
	var client = DarkSky.NewAPIClient("TEST_API_KEY")
	if client == nil {
		t.Error("couldn't create a new API client")
	}
}

func TestAPIClientSetUnits(t *testing.T) {
	var client = DarkSky.NewAPIClient("TEST_API_KEY")
	client.Units = DarkSky.SI
	if client.Units != DarkSky.SI {
		t.Error("couldn't set units on API client")
	}
}

func TestGetForecast(t *testing.T) {
	var client = DarkSky.NewAPIClient(os.Getenv("FORECAST_IO_API_KEY"))
	_, err := client.GetForecast(0, 0)
	if err != nil {
		t.Error(err)
	}
}
