package trimet

type Direction struct {
	// The number of the direction, either 1 for inbound or 0 for outbound.
	Number int `json:"dir"`

	// Describes the direction of the route.
	Description string `json:"desc"`

	// List of stops included in the direction of a route.
	Locations []Location `json:"stop"`
}
