package trimet

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestDetoursService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/detours", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"appID":  testAppID,
			"json":   "true",
			"routes": "12",
		})
		b, err := ioutil.ReadFile("testdata/detours.json")
		if nil != err {
			t.Fatal("Unable to read testdata/detours.json")
		}
		w.Write(b)
	})

	req := &DetoursRequest{
		Routes: []int{12},
	}
	detours, err := client.Detours.Get(req)

	if err != nil {
		t.Fatalf("Detours.Get returned error: %v", err)
	}

	PST, _ := time.LoadLocation("America/Los_Angeles")
	dt20131108140700 := time.Date(2013, 11, 8, 14, 7, 0, 0, PST)
	dt20371109020000 := time.Date(2037, 11, 9, 2, 0, 0, 0, PST)
	dt20131202030000 := time.Date(2013, 12, 2, 3, 0, 0, 0, PST)
	dt20371203020000 := time.Date(2037, 12, 3, 2, 0, 0, 0, PST)
	dt20140113144700 := time.Date(2014, 1, 13, 14, 47, 0, 0, PST)
	dt20370925020000 := time.Date(2037, 9, 25, 2, 0, 0, 0, PST)
	expect := &DetoursResponse{
		Detours: []Detour{
            {
                ID: "28997",
                Phonetic: "No service to SW Pacific Highway & 78th due to construction. Use stops before or after.",
                Description: "No service to SW Pacific Hwy & 78th (Stop ID 4305) due to construction. Use stops before or after.",
                Begin: &Time{&dt20131108140700},
                End: &Time{&dt20371109020000},
                Routes: []Route{
                    {
                        Detour: true,
                        Description: "12-Barbur/Sandy Blvd",
                        ID: 12,
                        Type: "B",
                    },
                },
            },
            {
                ID: "29416",
                Phonetic: "For trips to Portland City Center, no service to SW Barbur at Luradel due to construction. Use next stop at Huber St .",
                Description: "For trips to Portland City Center, no service to SW Barbur at Luradel due to construction. Use next stop at Huber St (Stop ID 150).",
                Begin: &Time{&dt20131202030000},
                End: &Time{&dt20371203020000},
                Routes: []Route{
                    {
                        Detour: true,
                        Description: "12-Barbur/Sandy Blvd",
                        ID: 12,
                        Type: "B",
                    },
                },
            },
            {
                ID: "29755",
                Phonetic: "The southbound stop on SW Barbur at Capitol Hill Rd. is closed. Use stop at Evans or a temporary stop at 21st.",
                Description: "The southbound stop on SW Barbur at Capitol Hill Rd is closed. Use stop at Evans (Stop ID 201) or a temporary stop at 21st.",
                Begin: &Time{&dt20140113144700},
                End: &Time{&dt20370925020000},
                Routes: []Route{
                    {
                        Detour: true,
                        Description: "12-Barbur/Sandy Blvd",
                        ID: 12,
                        Type: "B",
                    },
                },
            },
		},
	}
	if !reflect.DeepEqual(detours, expect) {
		t.Errorf("Expected Detours.Get to return:\n%+v\nfound:\n%+v", expect, detours)
	}
}
