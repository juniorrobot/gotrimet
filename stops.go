package trimet

// StopsService returns stops that are within a geographically defined area or
// within a distance of a point.
//
// TriMet API docs: http://developer.trimet.org/ws_docs/stop_location_ws.shtml
type StopsService struct {
	client *Client
}

type StopsRequest struct {
	Request

	// Define the lower left and upper right corners of the bounding box.
	//
	// Comma delimited list of longitude and latitude values.
	// Arguments are lonmin, latmin, lonmax, latmax in decimal degrees.
	BoundingBox struct {
		LonMin float64 `url:"lonmin"`
		LatMin float64 `url:"latmin"`
		LonMax float64 `url:"lonmax"`
		LatMax float64 `url:"latmax"`
	} `url:"bbox,omitempty"`

	// Defines center of search radius in decimal degrees.
	// Longitude, Latitude pair.
	LonLat []float64 `url:"ll,omitempty,comma"`

	// Use with LonLat to define search radius in feet.
	Feet int `url:"feet,omitempty"`

	// Use with LonLat to define search radius in meters.
	Meters int `url:"meters,omitempty"`

	// Whether to include a list of routes that service the stop(s).
	ShowRoutes bool `url:"showRoutes,omitempty"`

	// Whether to include a list of Direction elements for each route
	// direction that service the stop(s). Setting ShowRoutes to 'true' is
	// unnecessary if this is set to 'true'.
	ShowRouteDirections bool `url:"showRouteDirs,omitempty"`
}

type StopsResponse struct {
	Response
	Locations []Location `json:"location"`
}

type stopsResponseResults struct {
	Results *StopsResponse `json:"resultSet,omitempty"`
}

// Get latest stop information.
func (s *StopsService) Get(r *StopsRequest) (*StopsResponse, error) {
	response := new(stopsResponseResults)
	err := s.client.Get("stops", r, response)
	if nil != err {
		return nil, err
	}

	return response.Results, nil
}
