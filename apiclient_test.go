package darksky_test

import (
	"os"
	"testing"

	"github.com/sxg/DarkSky"
)

func TestNewAPIClient(t *testing.T) {
	var client = darksky.NewAPIClient("TEST_API_KEY")
	if client == nil {
		t.Error("couldn't create a new API client")
	}
}

func TestAPIClientSetUnits(t *testing.T) {
	var client = darksky.NewAPIClient("TEST_API_KEY")
	client.Units = darksky.SI
	if client.Units != darksky.SI {
		t.Error("couldn't set units on API client")
	}
}

func TestAPIClientSetLanguage(t *testing.T) {
	var client = darksky.NewAPIClient("TEST_API_KEY")
	client.Language = darksky.French
	if client.Language != darksky.French {
		t.Error("couldn't set language on API client")
	}
}

func TestGetForecast(t *testing.T) {
	var client = darksky.NewAPIClient(os.Getenv("FORECAST_IO_API_KEY"))
	_, err := client.GetForecast(0, 0)
	if err != nil {
		t.Error(err)
	}
}
