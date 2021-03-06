package darksky_test

import (
	"os"
	"testing"
	"time"

	"github.com/sxg/DarkSky"
)

func TestNewAPIClient(t *testing.T) {
	var client = darksky.NewAPIClient(os.Getenv("FORECAST_IO_API_KEY"))
	if client == nil {
		t.Error("couldn't create a new API client")
	}
}

func TestAPIClientSetUnits(t *testing.T) {
	var client = darksky.NewAPIClient(os.Getenv("FORECAST_IO_API_KEY"))
	client.Units = darksky.UnitsSI
	if client.Units != darksky.UnitsSI {
		t.Error("couldn't set units on API client")
	}
}

func TestAPIClientSetLanguage(t *testing.T) {
	var client = darksky.NewAPIClient(os.Getenv("FORECAST_IO_API_KEY"))
	client.Language = darksky.LanguageFrench
	if client.Language != darksky.LanguageFrench {
		t.Error("couldn't set language on API client")
	}
}

func TestGetForecast(t *testing.T) {
	var client = darksky.NewAPIClient(os.Getenv("FORECAST_IO_API_KEY"))
	_, err := client.GetForecast(30, 30)
	if err != nil {
		t.Error(err)
	}
}

func TestGetForecastAtTime(t *testing.T) {
	var client = darksky.NewAPIClient(os.Getenv("FORECAST_IO_API_KEY"))
	_, err := client.GetForecastAtTime(30, 30, time.Now())
	if err != nil {
		t.Error(err)
	}
}
