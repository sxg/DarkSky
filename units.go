package darksky

// Units in which the Dark Sky API data will be returned.
type Units string

const (
	// SI units.
	SI Units = "si"

	// US uses United States units and the default option.
	US = "us"

	// CA uses Canadian units.
	CA = "ca"

	// UK uses United Kingdom units.
	UK = "uk2"

	// Auto uses the appropriate units based on the location for which weather data is requested.
	Auto = "auto"
)
