package trimet

import (
	"time"
)

// Arrival contains arrival details for a Location.
type Arrival struct {
	// The Location id of the arrival.
	Location int `json:"locid"`

	// The block of the arrival.
	Block int `json:"block"`

	// The route number of the arrival.
	Route int `json:"route"`

	// Indicates if the vehicle has begun the trip which will arrive at the
	// indicated stop.
	Departed bool `json:"departed"`

	// Indicates if the arrival may be effected by a detour in effect along the
	// route.
	Detour bool `json:"detour"`

	// The direction of the route for this arrival.
	Direction int `json:"dir"`

	// Current status of the service.
	//
	// There are four possible values:
	//     estimated: Arrival time was estimated with vehicle position
	//       information
	//     scheduled: Scheduled arrival time is available only. No real
	//       time information available for estimation. Bus' radio may be
	//       down or vehicle may not be in service. Arrivals are not estimated
	//       when further than an hour away.
	//     delayed: Status of service is uncertain.
	//     canceled: Scheduled arrival was canceled for the day.
	Status string `json:"status"`

	// The estimated time for this arrival. If this value is not present the
	// arrival could not be estimated and schedule is shown instead.
	Estimated *time.Time `json:"estimated"`

	// The scheduled stop time (or interpolated scheduled stop time when the
	// stop is not a time point) of the arrival.
	Scheduled *time.Time `json:"scheduled"`

	// The full text of the overhead sign of the vehicle when it arrives at the
	// stop.
	FullSign string `json:"fullsign"`

	// The short version of text from the overhead sign of the vehicle when it
	// arrives at the stop.
	ShortSign string `json:"shortsign"`

	// The piece of the block for this arrival.
	Piece string `json:"piece"`

	// The last known position of the vehicle along its block. Includes path
	// information from this position to the indicated stop.
	BlockPosition Position `json:"blockPosition"`

	// Indicates conditions are influencing the reporting of arrivals for a
	// route. This occurs in inclement weather conditions.
	RouteStatus struct {
		// Route number of this status.
		Route int `json:"route"`

		// The most current reported status.
		//
		// Possible values:
		//     estimatedOnly: Arrivals for this route are only being reported
		//       if they can be estimated within the next hour. This occurs in
		//       inclement weather conditions.
		//     off: No arrivals are being reported for this route. This occurs
		//       when conditions such as snow and ice cause vehicles along the
		//       route to travel off their trip patterns. In such cases
		//       predictions are highly inaccurate or impossible.
		Status string `json:"status"`
	} `json:"routeStatus"`
}
