package darksky

// Icon to represent weather conditions. Additional values may be defined in the future, so be sure to use a default.
type Icon string

const (
	// IconClearDay represents a clear day.
	IconClearDay Icon = "clear-day"

	// IconClearNight represents a clear night.
	IconClearNight = "clear-night"

	// IconRain represents a rainy day or night.
	IconRain = "rain"

	// IconSnow represents a snowy day or night.
	IconSnow = "snow"

	// IconSleet represents a sleety day or night.
	IconSleet = "sleet"

	// IconWind represents a windy day or night.
	IconWind = "wind"

	// IconFog represents a foggy day or night.
	IconFog = "fog"

	// IconCloudy represents a cloudy day or night.
	IconCloudy = "cloudy"

	// IconPartlyCloudyDay represents a partly cloudy day.
	IconPartlyCloudyDay = "partly-cloudy-day"

	// IconPartlyCloudyNight represents a partly cloudy night.
	IconPartlyCloudyNight = "partly-cloudy-night"
)
