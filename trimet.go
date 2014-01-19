package trimet

import (
	"encoding/json"
	"errors"
	"fmt"
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

	// Application ID authorized by TriMet for API requests.
	appID string

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
func NewClient(appID string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client: httpClient,
		appID: appID,
		BaseURL: baseURL,
		UserAgent: userAgent,
	}
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
// params is included with the request query.
func (c *Client) NewRequest(method, urlStr string, params interface{}) (*http.Request, error) {
	if "" == c.appID {
		return nil, errors.New("Missing required AppID")
	}
	if "" == urlStr {
		return nil, errors.New("Requested URL must not be empty")
	}

	req := newRequest(c.appID)
	queryVals, err := query.Values(req)
	if nil != err {
		return nil, err
	}

	paramVals, err := parameterValues(params)
	if nil != err {
		return nil, err
	} else if nil != paramVals {
		for k, vals := range paramVals {
			for _, v := range vals {
				queryVals.Add(k, v)
			}
		}
	}

	rel, err := urlWithQuery(urlStr, queryVals)
	if nil != err {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	httpReq, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Accept", mediaType)
	httpReq.Header.Add("User-Agent", c.UserAgent)
	return httpReq, nil
}

// addParameters adds the parameters in params as URL query parameters to base.
// params must be a struct whose fields may contain "url" tags.
func parameterValues(params interface{}) (url.Values, error) {
	if nil == params {
		return nil, nil
	}

	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil, nil
	}

	return query.Values(params)
}

// urlWithQuery adds the parameters in params as URL query parameters to base.
// params must be a struct whose fields may contain "url" tags.
func urlWithQuery(base string, q url.Values) (*url.URL, error) {
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	u.RawQuery = q.Encode()
	return u, nil
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

	data, err := ioutil.ReadAll(response.Body)
	if nil == err && nil != data {
		err = CheckResponse(response, data)
		if nil == err && nil != v {
			err = json.Unmarshal(data, v)
		}
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
func CheckResponse(r *http.Response, data []byte) error {
	errorResponse := newErrorResponse(r)
	err := json.Unmarshal(data, errorResponse)
	if nil == err && "" != errorResponse.Message.Content {
		return errorResponse
	}

	if c := r.StatusCode; c < 200 || c > 299 {
		return fmt.Errorf("Encountered HTTP error status %v", c)
	}

	return nil
}

// Do sends an API request and returns the API response.
//
// The API response is decoded and stored in the value pointed to by v, or
// returned as an error if an API error has occurred.
func (c *Client) Get(url string, request interface{}, response interface{}) error {
	if nil == response {
		return errors.New("GET expects response data")
	}

	req, err := c.NewRequest("GET", url, request)
	if err != nil {
		return err
	}

	_, err = c.Do(req, response)
	if err != nil {
		return err
	}

	return nil
}
