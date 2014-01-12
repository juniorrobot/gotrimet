package trimet

// DetoursService retrieves a list of detours currently in effect by route.
//
// TriMet API docs: http://developer.trimet.org/ws_docs/detours_ws.shtml
type DetoursService struct {
}

type DetoursRequest struct {
	Request

	// If present results will contain only detours applicable for the route
	// numbers provided. If ommitted every detour in effect will be returned.
	Routes []string `url:"routes"`
}

type DetoursResponse struct {
	Response
	Detours []Detour `json:"detour"`
}
