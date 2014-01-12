package trimet

// RoutesService retrieves a list of routes being reported by TransitTracker
// from the active schedule, optionally a list of directions for those routes
// and stops in each of those directions.
//
// TriMet API docs: http://developer.trimet.org/ws_docs/routeConfig_ws.shtml
type RoutesService struct {
}

type RouteConfigRequest struct {
	Request

	// Include only the routes with these numbers.
	// If omitted every route will be returned.
	Routes []int `url:"routes,omitempty,comma"`

	// direction elements to include under route number.
	// Must be one of:
	//     0: outbound
	//     1: inbound
	//     'true' or 'yes': both directions
	Direction string `url:"dir,omitempty"`

	// If this argument is present and has any non-empty value, stop elements
	// will be included under each route direction element.
	Stops string `url:"stops,omitempty"`

	// If this argument is present and has any non-empty value, stop elements
	// will be included under each route direction element that are also time
	// points along the route. If this argument is used there is no need for
	// the Stops argument.
	TimePoints string `url:"tp,omitempty"`

	// Only stops with sequence numbers higher or equal to this value will
	// be included in stop lists.
	StartSequence int `url:"startSeq,omitempty"`

	// Only stops with sequence numbers lower or equal to this value will be
	// included in stop lists.
	EndSequence int `url:"endSeq,omitempty"`
}

type RouteConfigResponse struct {
	Response
	Routes []Route `json:"route"`
}
