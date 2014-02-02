package trimet

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestStopsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/stops", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"appID":         testAppID,
			"json":          "true",
			"ll":            "45.53055,-122.686153",
			"feet":          "500",
			"showRouteDirs": "true",
		})
		b, err := ioutil.ReadFile("testdata/stops.json")
		if nil != err {
			t.Fatal("Unable to read testdata/stops.json")
		}
		w.Write(b)
	})

	req := &StopsRequest{
		LonLat:              []float64{45.5305500, -122.6861530},
		Feet:                500,
		ShowRouteDirections: true,
	}
	stops, err := client.Stops.Get(req)

	if err != nil {
		t.Errorf("Stops.Get returned error: %v", err)
	}

	expect := &StopsResponse{
		Response: Response{QueryTime: newTestTime(t, "2014-01-12T15:32:13.438-0800")},
		Locations: []Location{
			{
				ID:          10775,
				Description: "NW Northrup & 14th",
				Direction:   "Westbound",
				Lon:         -122.685356502158,
				Lat:         45.5315030383606,
				Routes: []Route{
					{
						ID:          193,
						Type:        "R",
						Description: "Portland Streetcar - NS Line",
						Directions: []Direction{
							{
								Number:      0,
								Description: "To NW 23rd and Marshall",
							},
						},
					},
				},
			},
			{
				ID:          10752,
				Description: "NW Lovejoy & 13th",
				Direction:   "Eastbound",
				Lon:         -122.684611,
				Lat:         45.529997,
				Routes: []Route{
					{
						ID:          193,
						Type:        "R",
						Description: "Portland Streetcar - NS Line",
						Directions: []Direction{
							{
								Number:      1,
								Description: "To South Waterfront",
							},
						},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(stops, expect) {
		t.Errorf("Expected Stops.Get to return %+v, found %+v", expect, stops)
	}
}

func TestStopsService_Get_badRequest(t *testing.T) {
	var req *StopsRequest
	_, err := client.Stops.Get(req)
	if nil == err {
		t.Error("Expected Stops.Get to return error for nil request")
	}
}

func TestNewStopsWithCoords(t *testing.T) {
	var lat, lon float64 = 45.52414929707939, -122.68547059502453
	req := NewStopsRequestWithCoords(lat, lon)
	if nil == req {
		t.Fatal("Expected request to be created; received nil")
	}

	if nil == req.LonLat || 0 == len(req.LonLat) {
		t.Fatal("Unexpected empty lon/lat slice")
	}
	if req.LonLat[0] != lon {
		t.Errorf("Expected Longitude=%v, found %v", lon, req.LonLat[0])
	}
	if req.LonLat[1] != lat {
		t.Errorf("Expected Latitude=%v, found %v", lat, req.LonLat[1])
	}
}
