package trimet

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
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

	expect := &DetoursResponse{
		Detours: []Detour{
            {
                ID: "28997",
                Phonetic: "No service to SW Pacific Highway & 78th due to construction. Use stops before or after.",
                Description: "No service to SW Pacific Hwy & 78th (Stop ID 4305) due to construction. Use stops before or after.",
                Begin: newTestTime(t, "2013-11-08T14:07:00.000-0800"),
                End: newTestTime(t, "2037-11-09T02:00:00.000-0800"),
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
                Begin: newTestTime(t, "2013-12-02T03:00:00.000-0800"),
                End: newTestTime(t, "2037-12-03T02:00:00.000-0800"),
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
                Begin: newTestTime(t, "2014-01-13T14:47:00.000-0800"),
                End: newTestTime(t, "2037-09-25T02:00:00.000-0700"),
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

func TestDetoursService_Get_badRequest(t *testing.T) {
	var req *DetoursRequest
	_, err := client.Detours.Get(req)
	if nil == err {
		t.Error("Expected Detours.Get to return error for nil request")
	}
}
