package darksky

// Precipitation describes the type of precipitation.
type Precipitation string

const (
	// PrecipitationRain represents rainy precipitation.
	PrecipitationRain Precipitation = "rain"

	// PrecipitationSnow represents snowy precipitation.
	PrecipitationSnow = "snow"

	// PrecipitationSleet represents sleety precipitation.
	PrecipitationSleet = "sleet"

	// PrecipitationHail represents haily precipitation.
	PrecipitationHail = "hail"
)
