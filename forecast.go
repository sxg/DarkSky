package DarkSky

// Forecast weather data for a location at a specific time.
type Forecast struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Currently DataPoint `json:"currently"`
	Minutely  DataBlock `json:"minutely"`
	Hourly    DataBlock `json:"hourly"`
	Daily     DataBlock `json:"daily"`
	Alerts    []*Alert  `json:"alerts"`
	Flags     Flag      `json:"flags"`
}
