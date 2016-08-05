package darksky

// Units in which the Dark Sky API data will be returned.
type Units string

const (
	// UnitsSI units.
	UnitsSI Units = "si"

	// UnitsUS uses United States units and the default option.
	UnitsUS = "us"

	// UnitsCA uses Canadian units.
	UnitsCA = "ca"

	// UnitsUK uses United Kingdom units.
	UnitsUK = "uk2"

	// UnitsAuto uses the appropriate units based on the location for which weather data is requested.
	UnitsAuto = "auto"
)
