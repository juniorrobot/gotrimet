package trimet

import (
	"testing"
	"time"
)

func TestNewTime(t *testing.T) {
	PST, _ := time.LoadLocation("America/Los_Angeles")
	dt20140119120000 := time.Date(2014, 01, 19, 12, 0, 0, 0, PST)

	newTime := NewTime(dt20140119120000)
	if dt20140119120000 != *newTime.Time {
		t.Errorf("Expected new time to equal \"%v\", found \"%v\"", dt20140119120000, newTime)
	}
}

func checkParsedTime(t *testing.T, expect, actual time.Time) {
	expectYear, expectMonth, expectDay := expect.Date()
	expectHour, expectMin, expectSec := expect.Clock()
	expectTZ, _ := expect.Zone()
	actualYear, actualMonth, actualDay := actual.Date()
	actualHour, actualMin, actualSec := actual.Clock()
	actualTZ, _ := actual.Zone()
	if expectYear != actualYear ||
		expectMonth != actualMonth ||
		expectDay != actualDay ||
		expectHour != actualHour ||
		expectMin != actualMin ||
		expectSec != actualSec ||
		expectTZ != actualTZ {
		t.Errorf("Expected parsed time to equal \"%+v\", found \"%+v\"", expect, actual)
	}
}

func TestParseTime(t *testing.T) {
	PST, _ := time.LoadLocation("America/Los_Angeles")
	dt20140119120000 := time.Date(2014, 01, 19, 12, 0, 0, 0, PST)

	newTime, err := ParseTime("2014-01-19T12:00:00.000-0800")
	if nil != err {
		t.Fatalf("Unexpected error from ParseTime: %v", err)
	}

	checkParsedTime(t, dt20140119120000, *newTime.Time)
}

func TestParseTime_badFormat(t *testing.T) {
	_, err := ParseTime("2014-01-19T12:00:00.000Z0800")
	if nil == err {
		t.Fatal("Expected error to be returned from ParseTime for RFC3339Nano format")
	}
}

func TestUnmarshalTime(t *testing.T) {
	timestamp := `"2014-01-19T12:00:00.000-0800"`
	newTime := new(Time)
	err := newTime.UnmarshalJSON([]byte(timestamp))
	if nil != err {
		t.Fatalf("Unexpected error unmarshaling time %v: %v", timestamp, err)
	}

	PST, _ := time.LoadLocation("America/Los_Angeles")
	dt20140119120000 := time.Date(2014, 01, 19, 12, 0, 0, 0, PST)
	checkParsedTime(t, dt20140119120000, *newTime.Time)
}
