package trimet

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestArrivalsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/arrivals", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"appID":     testAppID,
			"json":      "true",
			"locIDs":    "8989",
			"streetcar": "true",
		})
		b, err := ioutil.ReadFile("testdata/arrivals.json")
		if nil != err {
			t.Fatal("Unable to read testdata/arrivals.json")
		}
		w.Write(b)
	})

	req := &ArrivalsRequest{LocationIDs: []int{8989}, Streetcar: true}
	arrivals, err := client.Arrivals.Get(req)

	if err != nil {
		t.Errorf("Arrivals.Get returned error: %v", err)
	}

	PST, _ := time.LoadLocation("America/Los_Angeles")
	dt20140112171209 := time.Date(2014, 01, 12, 17, 12, 9, 351000000, PST)
	dt20140112174600 := time.Date(2014, 01, 12, 17, 46, 0, 0, PST)
	dt20140112171205 := time.Date(2014, 01, 12, 17, 12, 5, 0, PST)
	expect := &ArrivalsResponse{
		Response: Response{
			QueryTime: &Time{&dt20140112171209},
		},
		Locations: []Location{
			{
				ID:          int(8989),
				Description: "NW 23rd & Marshall",
				Direction:   "Southbound",
				Lon:         -122.698688376761,
				Lat:         45.5306116478909,
			},
		},
		Arrivals: []Arrival{
			{
				Detour:    true,
				Status:    "estimated",
				Location:  8989,
				Block:     1537,
				Scheduled: &Time{&dt20140112174600},
				ShortSign: "15 Gateway TC",
				Direction: 1,
				Estimated: &Time{&dt20140112174600},
				Route:     15,
				Departed:  false,
				BlockPosition: Position{
					At:      &Time{&dt20140112171205},
					Feet:    15005,
					Lon:     -122.6973469,
					Lat:     45.5233678,
					Heading: 273,
					Trips: []Trip{
						{
							ID:          "4285706",
							Progress:    50383,
							Description: "Montgomery Park",
							Pattern:     21,
							Direction:   0,
							Route:       15,
							Distance:    60942,
						},
						{
							ID:          "4285964",
							Progress:    0,
							Description: "Gateway Layover",
							Pattern:     26,
							Direction:   1,
							Route:       15,
							Distance:    4447,
						},
					},
				},
				FullSign: "15  Belmont/NW 23rd to Gateway TC",
				Piece:    "1",
			},
		},
	}
	if !reflect.DeepEqual(arrivals, expect) {
		t.Errorf("Expected Arrivals.Get to return %+v, found %+v", expect, arrivals)
	}
}
