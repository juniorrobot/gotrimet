package trimet

// A Location contains information about a location, usually a stop.
//
// This element can occur in two locations in a response: one to
// describe the stop requested, and others to describe the location of any
// layovers.
type Location struct {
	ID int `json:"locid"`

	// The public location description of the stop.
	Description string `json:"desc"`

	// The direction of traffic at the stop.
	Direction string `json:"dir"`

	// The latitude of the stop.
	Lat float64 `json:"lat"`
	// The longitude of the stop.
	Lon float64 `json:"lng"`

	// List of routes that service the stop.
	Routes []Route `json:"route"`
}
