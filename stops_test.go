package trimet

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
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

	PST, _ := time.LoadLocation("America/Los_Angeles")
	dt20140112153213 := time.Date(2014, 01, 12, 15, 32, 13, 438000000, PST)
	expect := &StopsResponse{
		Response: Response{QueryTime: &Time{&dt20140112153213}},
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
