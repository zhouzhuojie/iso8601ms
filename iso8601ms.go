// Package iso8601ms package encodes and decodes time.Time in JSON in
// ISO 8601 with millisecond precision format
package iso8601ms

import (
	"time"
)

const (
	iso8601msFormat = "2006-01-02T15:04:05.000Z07:00"
)

// Time is a iso8601ms struct
type Time time.Time

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	s := `"` + tt.Format(iso8601msFormat) + `"`
	return []byte(s), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	var tt time.Time
	err := tt.UnmarshalJSON(data)
	*t = Time(tt)
	return err
}
