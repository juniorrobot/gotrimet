package trimet

import (
	"strings"
	"time"
)

// Time is a wrapper for time.Time to workaround various timezone issues.
//
// Times returned from TriMet are not in the proper RFC3339 format, as they
// lack the "Z" specifier separating seconds from timezone.
// Even if TriMet formats were correct, there are parsing bugs in Go that lead
// to an inability to parse a proper RFC3339 string with a timezone.
type Time struct {
	*time.Time
}

const trimetTime = "2006-01-02T15:04:05.999999999"// NB: no timezone

// UnmarshalJSON parses a TriMet time into a time.Time.
//
// It always assumes a PST timezone.
func (t *Time) UnmarshalJSON(data []byte) error {
	loc, err := time.LoadLocation("America/Los_Angeles")
	if nil != err {
		loc = time.FixedZone("PST", -8*60*60)// FIXME: handle DST
	}

	r := strings.NewReplacer("-0800", "", "-0700", "")
	fixed := r.Replace(string(data))
	parsed, err := time.ParseInLocation(`"`+trimetTime+`"`, fixed, loc)
	if nil == err {
		t.Time = new(time.Time)
		*t.Time = parsed
	}
	return err
}
