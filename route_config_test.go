package trimet

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestRoutesService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/routeConfig", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"appID":  testAppID,
			"json":   "true",
			"routes": "193",
			"tp":     "true",
		})
		b, err := ioutil.ReadFile("testdata/routeConfig.json")
		if nil != err {
			t.Fatal("Unable to read testdata/routeConfig.json")
		}
		w.Write(b)
	})

	req := &RouteConfigRequest{
		Routes:     []int{193},
		TimePoints: "true",
	}
	routes, err := client.Routes.Get(req)

	if err != nil {
		t.Fatalf("Routes.Get returned error: %v", err)
	}

	expect := &RouteConfigResponse{
		Routes: []Route{
			{
				ID:          193,
				Type:        "R",
				Description: "Portland Streetcar - NS Line",
				Directions: []Direction{
					{
						Number:      0,
						Description: "To NW 23rd and Marshall",
						Locations: []Location{
							{
								Description: "SW Lowell & Bond",
								ID:          12881,
								TimePoint:   true,
								Sequence:    25,
								Lon:         -122.671376020374,
								Lat:         45.4938906298509,
							},
							{
								Description: "SW Bond & Lane",
								ID:          12882,
								TimePoint:   false,
								Sequence:    50,
								Lon:         -122.670932716808,
								Lat:         45.495593953864,
							},
							{
								Description: "OHSU Commons",
								ID:          12883,
								TimePoint:   true,
								Sequence:    100,
								Lon:         -122.670738623655,
								Lat:         45.4989385765801,
							},
							{
								Description: "SW Moody & Meade",
								ID:          13602,
								TimePoint:   false,
								Sequence:    150,
								Lon:         -122.672742267264,
								Lat:         45.5033040096853,
							},
							{
								Description: "SW River Pkwy & Moody",
								ID:          12379,
								TimePoint:   true,
								Sequence:    200,
								Lon:         -122.674139972701,
								Lat:         45.5071394700201,
							},
							{
								Description: "SW Harrison Street",
								ID:          12380,
								TimePoint:   false,
								Sequence:    250,
								Lon:         -122.676531233424,
								Lat:         45.5089495445265,
							},
							{
								Description: "SW 1st & Harrison",
								ID:          12381,
								TimePoint:   false,
								Sequence:    300,
								Lon:         -122.677878143433,
								Lat:         45.5097608385749,
							},
							{
								Description: "SW 3rd & Harrison",
								ID:          12382,
								TimePoint:   false,
								Sequence:    350,
								Lon:         -122.679813063405,
								Lat:         45.5102771993458,
							},
							{
								Description: "PSU Urban Center",
								ID:          10764,
								TimePoint:   true,
								Sequence:    400,
								Lon:         -122.682078,
								Lat:         45.51222,
							},
							{
								Description: "SW Park & Mill",
								ID:          10766,
								TimePoint:   false,
								Sequence:    450,
								Lon:         -122.684553,
								Lat:         45.513054,
							},
							{
								Description: "SW 10th & Clay",
								ID:          10765,
								TimePoint:   true,
								Sequence:    500,
								Lon:         -122.684978,
								Lat:         45.514546,
							},
							{
								Description: "Art Museum",
								ID:          6493,
								TimePoint:   false,
								Sequence:    550,
								Lon:         -122.68399099998,
								Lat:         45.516304999998,
							},
							{
								Description: "Central Library",
								ID:          10767,
								TimePoint:   true,
								Sequence:    600,
								Lon:         -122.682471,
								Lat:         45.519225,
							},
							{
								Description: "SW 10th & Alder",
								ID:          10768,
								TimePoint:   false,
								Sequence:    650,
								Lon:         -122.681733,
								Lat:         45.520573,
							},
							{
								Description: "SW 10th & Stark",
								ID:          10769,
								TimePoint:   false,
								Sequence:    700,
								Lon:         -122.681090913694,
								Lat:         45.5217417342333,
							},
							{
								Description: "NW 10th & Couch",
								ID:          10770,
								TimePoint:   false,
								Sequence:    750,
								Lon:         -122.681083,
								Lat:         45.523593,
							},
							{
								Description: "NW 10th & Everett",
								ID:          10771,
								TimePoint:   false,
								Sequence:    800,
								Lon:         -122.681113,
								Lat:         45.525011,
							},
							{
								Description: "NW 10th & Glisan",
								ID:          10772,
								TimePoint:   false,
								Sequence:    850,
								Lon:         -122.68118,
								Lat:         45.526446,
							},
							{
								Description: "NW 10th & Johnson",
								ID:          10773,
								TimePoint:   true,
								Sequence:    900,
								Lon:         -122.68125,
								Lat:         45.528572,
							},
							{
								Description: "NW 10th & Northrup",
								ID:          13604,
								TimePoint:   false,
								Sequence:    950,
								Lon:         -122.681365907158,
								Lat:         45.5314381810721,
							},
							{
								Description: "NW 12th & Northrup",
								ID:          12796,
								TimePoint:   false,
								Sequence:    1000,
								Lon:         -122.683319529015,
								Lat:         45.5315346845716,
							},
							{
								Description: "NW Northrup & 14th",
								ID:          10775,
								TimePoint:   true,
								Sequence:    1050,
								Lon:         -122.685356502158,
								Lat:         45.5315030383606,
							},
							{
								Description: "NW Northrup & 18th",
								ID:          10776,
								TimePoint:   true,
								Sequence:    1100,
								Lon:         -122.689416558363,
								Lat:         45.5314335086312,
							},
							{
								Description: "NW Northrup & 21st",
								ID:          10777,
								TimePoint:   false,
								Sequence:    1150,
								Lon:         -122.694455,
								Lat:         45.531346,
							},
							{
								Description: "NW Northrup & 22nd",
								ID:          10778,
								TimePoint:   false,
								Sequence:    1200,
								Lon:         -122.696445,
								Lat:         45.531308,
							},
							{
								Description: "NW 23rd & Marshall",
								ID:          8989,
								TimePoint:   true,
								Sequence:    1250,
								Lon:         -122.698688376761,
								Lat:         45.5306116478909,
							},
						},
					},
					{
						Number:      1,
						Description: "To South Waterfront",
						Locations: []Location{
							{
								Description: "NW 23rd & Marshall",
								ID:          8989,
								TimePoint:   true,
								Sequence:    50,
								Lon:         -122.698688376761,
								Lat:         45.5306116478909,
							},
							{
								Description: "NW Lovejoy & 22nd",
								ID:          3596,
								TimePoint:   false,
								Sequence:    100,
								Lon:         -122.69688,
								Lat:         45.529746,
							},
							{
								Description: "NW Lovejoy & 21st",
								ID:          3595,
								TimePoint:   false,
								Sequence:    150,
								Lon:         -122.694676019495,
								Lat:         45.5298329830986,
							},
							{
								Description: "NW Lovejoy & 18th",
								ID:          10751,
								TimePoint:   true,
								Sequence:    200,
								Lon:         -122.689587149344,
								Lat:         45.5299254165705,
							},
							{
								Description: "NW Lovejoy & 13th",
								ID:          10752,
								TimePoint:   true,
								Sequence:    250,
								Lon:         -122.684611,
								Lat:         45.529997,
							},
							{
								Description: "NW 11th & Johnson",
								ID:          10753,
								TimePoint:   true,
								Sequence:    300,
								Lon:         -122.682373998868,
								Lat:         45.5287417489584,
							},
							{
								Description: "NW 11th & Glisan",
								ID:          10754,
								TimePoint:   false,
								Sequence:    350,
								Lon:         -122.682297014895,
								Lat:         45.5266046660366,
							},
							{
								Description: "NW 11th & Everett",
								ID:          10755,
								TimePoint:   false,
								Sequence:    400,
								Lon:         -122.682245856996,
								Lat:         45.5251787559408,
							},
							{
								Description: "NW 11th & Couch",
								ID:          10756,
								TimePoint:   false,
								Sequence:    450,
								Lon:         -122.682223,
								Lat:         45.523784,
							},
							{
								Description: "SW 11th & Alder",
								ID:          9600,
								TimePoint:   true,
								Sequence:    500,
								Lon:         -122.68281899998,
								Lat:         45.521093999998,
							},
							{
								Description: "SW 11th & Taylor",
								ID:          9633,
								TimePoint:   false,
								Sequence:    550,
								Lon:         -122.683873318603,
								Lat:         45.5190589217565,
							},
							{
								Description: "SW 11th & Jefferson",
								ID:          10759,
								TimePoint:   false,
								Sequence:    600,
								Lon:         -122.685301013972,
								Lat:         45.5164024253733,
							},
							{
								Description: "SW 11th & Clay",
								ID:          10760,
								TimePoint:   false,
								Sequence:    650,
								Lon:         -122.686081,
								Lat:         45.515106,
							},
							{
								Description: "SW Park & Market",
								ID:          11011,
								TimePoint:   false,
								Sequence:    700,
								Lon:         -122.683913,
								Lat:         45.513704,
							},
							{
								Description: "SW 5th & Market",
								ID:          10762,
								TimePoint:   false,
								Sequence:    750,
								Lon:         -122.681041895921,
								Lat:         45.5129219831852,
							},
							{
								Description: "SW 5th & Montgomery",
								ID:          10763,
								TimePoint:   true,
								Sequence:    800,
								Lon:         -122.681314606923,
								Lat:         45.5117080762786,
							},
							{
								Description: "SW 3rd & Harrison",
								ID:          12375,
								TimePoint:   false,
								Sequence:    850,
								Lon:         -122.679597720418,
								Lat:         45.510203002331,
							},
							{
								Description: "SW 1st & Harrison",
								ID:          12376,
								TimePoint:   false,
								Sequence:    900,
								Lon:         -122.677679169304,
								Lat:         45.5096895341786,
							},
							{
								Description: "SW Harrison Street",
								ID:          12377,
								TimePoint:   false,
								Sequence:    950,
								Lon:         -122.676573625166,
								Lat:         45.5087917088502,
							},
							{
								Description: "SW River Pkwy & Moody",
								ID:          12378,
								TimePoint:   true,
								Sequence:    1000,
								Lon:         -122.673923439765,
								Lat:         45.5070639368801,
							},
							{
								Description: "SW Moody & Meade",
								ID:          13601,
								TimePoint:   false,
								Sequence:    1025,
								Lon:         -122.672759735753,
								Lat:         45.5030721811026,
							},
							{
								Description: "SW Moody & Gibbs",
								ID:          12760,
								TimePoint:   false,
								Sequence:    1050,
								Lon:         -122.671814843588,
								Lat:         45.4993370063905,
							},
							{
								Description: "SW Moody & Gaines",
								ID:          12880,
								TimePoint:   false,
								Sequence:    1100,
								Lon:         -122.671942044256,
								Lat:         45.4961824454633,
							},
							{
								Description: "SW Lowell & Bond",
								ID:          12881,
								TimePoint:   true,
								Sequence:    1150,
								Lon:         -122.671376020374,
								Lat:         45.4938906298509,
							},
						},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(routes, expect) {
		t.Errorf("Expected Routes.Get to return:\n%+v\nfound:\n%+v", expect, routes)
	}
}

func TestRoutesService_Get_badRequest(t *testing.T) {
	var req *RouteConfigRequest
	_, err := client.Routes.Get(req)
	if nil == err {
		t.Error("Expected Routes.Get to return error for nil request")
	}
}
