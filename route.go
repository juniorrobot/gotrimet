package trimet

type Route struct {
	// The route's number.
	ID int `json:"route"`

	// The route's description.
	Description string `json:"desc"`

	// The type of the route, either 'B' for bus, or 'R' for fixed guideway
	// (either rail or aerial tram).
	Type string `json:"type"`

	// Indicates if this route has a detour in effect.
	Detour bool `json:"detour"`

	// Information for each route direction.
	Directions []Direction `json:"dir"`
}
