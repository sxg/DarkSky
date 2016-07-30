package DarkSky

// Flag containing various metadata related to a Forecast request.
type Flag struct {
	DarkSkyUnavailable bool     `json:"darksky-unavailable"`
	DarkSkyStations    []string `json:"darksky-stations"`
	DataPointStations  []string `json:"datapoint-stations"`
	ISDStations        []string `json:"isd-stations"`
	LAMPStations       []string `json:"lamp-stations"`
	METARStations      []string `json:"metar-stations"`
	MetnoLicense       bool     `json:"metno-license"`
	Sources            []string `json:"sources"`
	Units              string   `json:"units"`
}
