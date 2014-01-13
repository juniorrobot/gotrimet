package trimet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the TriMet client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

const testAppID = `abc123`

// setup sets up a test HTTP server along with a trimet.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(testAppID, nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expect string) {
	if expect != r.Method {
		t.Errorf("Expected request method %v, found %v", expect, r.Method)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	expect := url.Values{}
	for k, v := range values {
		expect.Add(k, v)
	}

	r.ParseForm()
	if !reflect.DeepEqual(expect, r.Form) {
		t.Errorf("Expected request parameters %v, found %v", expect, r.Form)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, expect string) {
	if value := r.Header.Get(header); expect != value {
		t.Errorf("Expected Header \"%s = %s\", found: %s", header, expect, value)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, expect string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Unable to read body")
	}
	str := string(b)
	if expect != str {
		t.Errorf("Expected body = %s, found: %s", expect, str)
	}
}

func testJSONMarshal(t *testing.T, v interface{}, expect string) {
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", v)
	}

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(expect))
	if err != nil {
		t.Errorf("String is not valid json: %s", expect)
	}

	if w.String() != string(j) {
		t.Errorf("json.Marshal(%q) returned %s, expect %s", v, j, w)
	}

	u := reflect.ValueOf(v).Interface()
	if err := json.Unmarshal([]byte(expect), u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v", expect)
	}

	if !reflect.DeepEqual(v, u) {
		t.Errorf("json.Unmarshal(%q) returned %s, expect %s", expect, u, v)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(testAppID, nil)

	if testAppID != c.appID {
		t.Errorf("Expected NewClient AppID = %v, found %v", testAppID, c.appID)
	}
	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("Expected NewClient BaseURL = %v, found %v", defaultBaseURL, c.BaseURL.String())
	}
	if c.UserAgent != userAgent {
		t.Errorf("Expected NewClient UserAgent = %v, found %v", userAgent, c.UserAgent)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(testAppID, nil)

	inURL, outURL := "foo", defaultBaseURL+"foo?appID="+testAppID+"&json=true"
	req, err := c.NewRequest("GET", inURL, nil)
	if nil != err {
		t.Fatalf("Unexpected error creating request: %v", err)
	}

	// test that relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("Expected NewRequest(%v) URL = %v, found %v", inURL, outURL, req.URL)
	}

	// test that default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("Expected NewRequest() User-Agent = %v, found %v", c.UserAgent, userAgent)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(testAppID, nil)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Expected request method = %v, found %v", m, r.Method)
		}
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	client.Do(req, body)

	expect := &foo{"a"}
	if !reflect.DeepEqual(body, expect) {
		t.Errorf("Expected response body = %v, found %v", expect, body)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

func TestCheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader(`{"errorMessage":{"content":"m"}}`)),
	}
	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Errorf("Expected error response.")
	}

	expect := &ErrorResponse{http: res}
	expect.Message.Content = "m"
	if !reflect.DeepEqual(err, expect) {
		t.Errorf("Expected error = %#v, found %#v", expect, err)
	}
}

func TestErrorResponse_Error(t *testing.T) {
	res := &http.Response{Request: &http.Request{}}
	err := ErrorResponse{http: res}
	err.Message.Content = "m"
	if err.Error() == "" {
		t.Errorf("Expected non-empty ErrorResponse.Error()")
	}
}
