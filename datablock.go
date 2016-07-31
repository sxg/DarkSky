package darksky

// DataBlock contains weather data for a specific location over a period of time.
type DataBlock struct {
	Summary *string       `json:"summary"`
	Icon    *string       `json:"icon"`
	Data    *[]*DataPoint `json:"data"`
}
