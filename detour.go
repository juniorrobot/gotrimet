package trimet

import (
	"time"
)

// A Detour contains information about a detour that may apply to one or more
// routes at the time the query was made.
type Detour struct {
	// A unique identifier of the detour.
	ID int `json:"id"`

	// Time the detour begins. This will always be a time in the past.
	// This field is used internally and may be of little use
	// outside of TriMet.
	Begin *time.Time `json:"begin"`

	// The time the detour will become invalid. Note that this will always be a
	// time in the future. Some end times will be very far in the future and
	// will be removed once the detour is no longer in effect. This field is
	// used internally and may be of little use outside of TriMet.
	End *time.Time `json:"end"`

	// A plain text description of the detour.
	Description string `json:"desc"`

	// A phonetic spelling of the route detour. This field is used by TriMet's
	// 238-Ride text-to-speech system.
	Phonetic string `json:"phonetic"`

	// Occurs for every route the detour is applicable.
	Route []Route `json:"route"`
}
