package trimet

import (
	"fmt"
	"net/http"
	"time"
)

// Response is a TriMet API response.
//
// This wraps the standard http.Response returned from TriMet and provides
// convenient access to things like query times.
type Response struct {
	QueryTime *time.Time `json:"queryTime"`
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response

	http *http.Response

	Message struct {
		Content string `json:"content"`
	} `json:"errorMessage"`
}

// newErrorResponse creates a new ErrorResponse for the provided http.Response.
func newErrorResponse(r *http.Response) *ErrorResponse {
	return &ErrorResponse{
		http: r,
	}
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.http.Request.Method, r.http.Request.URL,
		r.http.StatusCode, r.Message)
}
