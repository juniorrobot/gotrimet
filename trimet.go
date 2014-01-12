package trimet

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://developer.trimet.org/ws/V1/"
	userAgent      = "gotrimet/" + libraryVersion
	mediaType      = "application/json"
)

// A Client manages communication with the TriMet API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.  Defaults to the public TriMet API, but can be
	// set to a domain endpoint to use with beta features.  BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the TriMet API.
	UserAgent string

	// Services used for talking to different parts of the TriMet API.
	Arrivals *ArrivalsService
	Detours  *DetoursService
	Routes   *RoutesService
	Stops    *StopsService
}

// NewClient returns a new TriMet API client.
//
// If a nil httpClient is provided, http.DefaultClient will be used.  To use
// API methods which require authentication, provide an http.Client that will
// perform the authentication for you (such as that provided by the goauth2
// library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.Arrivals = &ArrivalsService{client: c}
	c.Detours = &DetoursService{client: c}
	c.Routes = &RoutesService{client: c}
	c.Stops = &StopsService{client: c}
	return c
}

// NewRequest creates an API request.
//
// A relative URL can be provided in urlStr, in which case it is resolved
// relative to the BaseURL of the Client.  Relative URLs should always be
// specified without a preceding slash.  If specified, the value pointed to by
// body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// addParameters adds the parameters in params as URL query parameters to base.
// params must be a struct whose fields may contain "url" tags.
func addParameters(base string, params interface{}) (string, error) {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return base, nil
	}

	u, err := url.Parse(base)
	if err != nil {
		return base, err
	}

	qs, err := query.Values(params)
	if err != nil {
		return base, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// Do sends an API request and returns the API response.
//
// The API response is decoded and stored in the value pointed to by v, or
// returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = CheckResponse(response)
	if err != nil {
		// even though there was an error, we still return the response in case
		// the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(response.Body).Decode(v)
	}
	return response, err
}

// CheckResponse checks the API response for errors, and returns them if
// present.
//
// A response is considered an error if it has a status code outside the 200
// range.  API error responses are expected to have either no response body, or
// a JSON response body that maps to ErrorResponse.  Any other response body
// will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := newErrorResponse(r)
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// Do sends an API request and returns the API response.
//
// The API response is decoded and stored in the value pointed to by v, or
// returned as an error if an API error has occurred.
func (c *Client) Get(url string, response interface{}) error {
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(req, response)
	if err != nil {
		return err
	}

	return nil
}
