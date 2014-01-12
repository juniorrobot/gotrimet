package trimet

type Request struct {
	// Authorized application ID.
	AppID string `url:"appID"`

	// If true results will be returned in JSON format rather than the default
	// XML format.
	JSON bool `url:"json,omitempty"`

	// If present returns the JSON result in a JSONP callback function. Only
	// used if JSON is set to true.
	Callback string `url:"callback,omitempty"`
}
