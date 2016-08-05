package darksky

// Icon to represent weather conditions. Additional values may be defined in the future, so be sure to use a default.
type Icon string

const (
	// ClearDay represents a clear day.
	ClearDay Icon = "clear-day"

	// ClearNight represents a clear night.
	ClearNight = "clear-night"

	// Rain represents a rainy day or night.
	Rain = "rain"

	// Snow represents a snowy day or night.
	Snow = "snow"

	// Sleet represents a sleety day or night.
	Sleet = "sleet"

	// Wind represents a windy day or night.
	Wind = "wind"

	// Fog represents a foggy day or night.
	Fog = "fog"

	// Cloudy represents a cloudy day or night.
	Cloudy = "cloudy"

	// PartlyCloudyDay represents a partly cloudy day.
	PartlyCloudyDay = "partly-cloudy-day"

	// PartlyCloudyNight represents a partly cloudy night.
	PartlyCloudyNight = "partly-cloudy-night"
)
