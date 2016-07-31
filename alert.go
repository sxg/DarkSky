package darksky

// Alert for a severe weather warning issued for a location by a governmental authority (consult the Dark Sky API documentation for a full list).
type Alert struct {
	Title       *string  `json:"title"`
	Expires     *float64 `json:"expires"`
	Description *string  `json:"description"`
	URI         *string  `json:"uri"`
}
