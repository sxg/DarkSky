package darksky

// DataPoint contains weather data for a specific location and time.
type DataPoint struct {
	Time                        *float64 `json:"time"`
	Summary                     *string  `json:"summary"`
	Icon                        Icon     `json:"icon"`
	SunriseTime                 *float64 `json:"sunriseTime"`
	SunsetTime                  *float64 `json:"sunsetTime"`
	MoonPhase                   *float64 `json:"moonPhase"`
	MoonPhaseError              *float64 `json:"moonPhaseErorr"`
	NearestStormDistance        *float64 `json:"nearestStormDistance"`
	NearestStormDistanceError   *float64 `json:"NearestStormDistanceError"`
	NearestStormBearing         *float64 `json:"nearestStormBearing"`
	NearestStormBearingError    *float64 `json:"nearestStormBearingError"`
	PrecipIntensity             *float64 `json:"precipIntensity"`
	PrecipIntensityError        *float64 `json:"precipIntensityError"`
	PrecipIntensityMax          *float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxError     *float64 `json:"precipIntensityMaxError"`
	PrecipIntensityMaxTime      *float64 `json:"precipIntensityMaxTime"`
	PrecipProbability           *float64 `json:"precipProbability"`
	PrecipProbabilityError      *float64 `json:"precipProbabilityError"`
	PrecipType                  *string  `json:"precipType"`
	PrecipAccumulation          *float64 `json:"precipAccumulation"`
	PrecipAccumulationError     *float64 `json:"precipAccumulationError"`
	Temperature                 *float64 `json:"temperature"`
	TemperatureError            *float64 `json:"temperatureError"`
	TemperatureMin              *float64 `json:"temperatureMin"`
	TemperatureMinError         *float64 `json:"temperatureMinError"`
	TemperatureMinTime          *float64 `json:"temperatureMinTime"`
	TemperatureMax              *float64 `json:"temperatureMax"`
	TemperatureMaxError         *float64 `json:"temperatureMaxError"`
	TemperatureMaxTime          *float64 `json:"temperatureMaxTime"`
	ApparentTemperature         *float64 `json:"apparentTemperature"`
	ApparentTemperatureError    *float64 `json:"apparentTemperatureError"`
	ApparentTemperatureMin      *float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinError *float64 `json:"apparentTemperatureMinError"`
	ApparentTemperatureMinTime  *float64 `json:"apparentTemperatureMinTime"`
	ApparentTemperatureMax      *float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxError *float64 `json:"apparentTemperatureMaxError"`
	ApparentTemperatureMaxTime  *float64 `json:"apparentTemperatureMaxTime"`
	DewPoint                    *float64 `json:"dewPoint"`
	DewPointError               *float64 `json:"dewPointError"`
	WindSpeed                   *float64 `json:"windSpeed"`
	WindSpeedError              *float64 `json:"windSpeedError"`
	WindBearing                 *float64 `json:"windBearing"`
	WindBearingError            *float64 `json:"windBearingError"`
	CloudCover                  *float64 `json:"cloudCover"`
	CloudCoverError             *float64 `json:"cloudCoverError"`
	Humidity                    *float64 `json:"humidity"`
	HumidityError               *float64 `json:"humidityError"`
	Pressure                    *float64 `json:"pressure"`
	PressureError               *float64 `json:"pressureError"`
	Visibility                  *float64 `json:"visibility"`
	VisibilityError             *float64 `json:"visibilityError"`
	Ozone                       *float64 `json:"ozone"`
	OzoneError                  *float64 `json:"ozoneError"`
}
