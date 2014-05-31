package trimet

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
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

	expect := &ArrivalsResponse{
		Response: Response{
			QueryTime: newTestTime(t, "2014-01-12T17:12:09.351-0800"),
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
				Scheduled: newTestTime(t, "2014-01-12T17:46:00.000-0800"),
				ShortSign: "15 Gateway TC",
				Direction: 1,
				Estimated: newTestTime(t, "2014-01-12T17:46:00.000-0800"),
				Route:     15,
				Departed:  false,
				BlockPosition: Position{
					At:      newTestTime(t, "2014-01-12T17:12:05.000-0800"),
					Feet:    15005,
					Lon:     -122.6973469,
					Lat:     45.5233678,
					Heading: 273,
					Trips: []Trip{
						{
							ID:          4285706,
							Progress:    50383,
							Description: "Montgomery Park",
							Pattern:     21,
							Direction:   0,
							Route:       15,
							Distance:    60942,
						},
						{
							ID:          4285964,
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
		t.Errorf("Expected Arrivals.Get to return:\nt%+v\nfound:\n%+v", expect, arrivals)
	}
}

func TestArrivalsService_Get_badRequest(t *testing.T) {
	var req *ArrivalsRequest
	_, err := client.Arrivals.Get(req)
	if nil == err {
		t.Error("Expected Arrivals.Get to return error for nil request")
	}
}
