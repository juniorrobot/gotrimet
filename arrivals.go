package trimet

// ArrivalsService reports next arrivals at a stop identified by location ID.
//
// TriMet API docs: http://developer.trimet.org/ws_docs/arrivals_ws.shtml
type ArrivalsService struct {
}

type ArrivalsRequest struct {
	Request

	// The location IDs for which to report arrivals.
	//
	// Arrivals are reported for each unique route and direction that services
	// each stop identified by their location ID. Up to 10 location IDs can be
	// reported at once.
	LocationIDs []int `url:"locIDs,comma"`

	// If true, NextBus API results will be included for those location IDs
	// served by Portland Streetcar.
	//
	// The results are transformed into arrival elements. If the NextBus server
	// is not responding or is unreachable, scheduled Portland Streetcar
	// scheduled times are provided:
	//     Status will be "estimated" for NextBus predictions.
	//     Status will be "scheduled" otherwise.
	//
	// Default is false.
	Streetcar bool `url:"streetcar,omitempty"`
}

type ArrivalsResponse struct {
	Response
	Locations []Location `json:"location"`
	Arrivals  []Arrival  `json:"arrival"`
}
