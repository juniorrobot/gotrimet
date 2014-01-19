package trimet

import (
	"time"
)

// Time is a wrapper for time.Time to workaround format issues.
//
// Times returned from TriMet are not in the proper RFC3339 format, as they
// lack the "Z" specifier separating seconds from timezone.
type Time struct {
	*time.Time
}

const trimetTime = `2006-01-02T15:04:05.999-0700`

// UnmarshalJSON parses a TriMet time into a time.Time.
func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(`"`+trimetTime+`"`, string(data))
	if nil == err {
		t.Time = new(time.Time)
		*t.Time = parsed
	}
	return err
}

// NewTime returns a new Time wrapping the given time.
func NewTime(t time.Time) *Time {
	return &Time{
		Time: &t,
	}
}

// ParseTime attempts to parse the timestamp using the TriMet format.
func ParseTime(timestamp string) (*Time, error) {
	parsed, err := time.Parse(trimetTime, string(timestamp))
	if nil != err {
		return nil, err
	}

	return NewTime(parsed), nil
}
