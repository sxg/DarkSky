package DarkSky_test

import (
	"fmt"
	"testing"

	"github.com/sxg/DarkSky"
)

func TestGetForecast(t *testing.T) {
	var client DarkSky.APIClient
	f, err := client.GetForecast(0, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f.Flags)
}
