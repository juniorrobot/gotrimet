package trimet

import (
	"time"
)

type Position struct {
	// The time this position was reported.
	At *time.Time `json:"at"`

	// Number of feet the vehicle is away from the stop at the time the
	// position was reported.
	Feet int `json:"feet"`

	// The heading of the vehicle at the time of the position was reported.
	Heading int `json:"heading"`

	// The latitude of the vehicle at the time the position was reported.
	Lat float64 `json:"lat"`

	// The longitude of the vehicle at the time the position was reported.
	Lon float64 `json:"lng"`

	// Occurs for every trip the vehicle must traverse to arrive at a stop.
	Trips []Trip `json:"trip"`

	// Occurs for every layover the vehicle has between its position and the
	// requested arrival.
	Layover struct {
		// The time the layover begins.
		Start *time.Time `json:"start"`

		// The time the layover ends.
		End *time.Time `json:"end"`
	} `json:"layover"`
}
