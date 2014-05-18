package trimet

type Trip struct {
	// The trip number of this trip.
	ID string `json:"tripNum"`

	// The route's direction description of the trip.
	Description string `json:"desc"`

	// The number of feet along a trip the vehicle must traverse to arrive at
	// a requested stop. If the vehicle must traverse the entire trip this
	// number will always be the entire length of the trip.
	Distance Distance `json:"destDist"`

	// The direction of the route of this trip.
	Direction int `json:"dir"`

	// The pattern number for the trip.
	Pattern int `json:"pattern"`

	// The number of feet the vehicle has traversed along a trip's pattern.
	Progress int `json:"progress"`

	// The route number for the related trip.
	Route int `json:"route"`
}
