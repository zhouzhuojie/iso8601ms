package iso8601ms

import (
	"encoding/json"
	"testing"
	"time"
)

var jsonTests = []struct {
	time time.Time
	json string
}{
	{time.Date(9999, 4, 12, 23, 20, 50, 520*1e6, time.UTC), `"9999-04-12T23:20:50.520Z"`},
	{time.Date(1996, 12, 19, 16, 39, 57, 0, time.FixedZone("UTC-4", -4*60*60)), `"1996-12-19T20:39:57.000Z"`},
	{time.Date(1996, 12, 19, 16, 39, 57, 1*1e6, time.FixedZone("UTC-7", -7*60*60)), `"1996-12-19T23:39:57.001Z"`},
	{time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), `"0000-01-01T00:00:00.000Z"`},
}

func TestNow(t *testing.T) {
	iso8601msTime := Time(time.Now())
	jsonBytes, err := json.Marshal(iso8601msTime)
	if err != nil {
		t.Errorf("error = %v, want nil", err)
	}
	if len(jsonBytes) != len(iso8601msFormat)+2 {
		t.Errorf("invalid json string length %v, want %v", len(jsonBytes), len(iso8601msFormat)+2)
	}
}

func TestIso8601msJSON(t *testing.T) {
	for _, tt := range jsonTests {
		iso8601msTime := Time(tt.time.UTC())
		var jsonIso8601msTime Time

		if jsonBytes, err := json.Marshal(iso8601msTime); err != nil {
			t.Errorf("%v json.Marshal error = %v, want nil", iso8601msTime, err)
		} else if string(jsonBytes) != tt.json {
			t.Errorf("%v JSON = %#q, want %#q", iso8601msTime, string(jsonBytes), tt.json)
		} else if err = json.Unmarshal(jsonBytes, &jsonIso8601msTime); err != nil {
			t.Errorf("%v json.Unmarshal error = %v, want nil", iso8601msTime, err)
		} else if !equalTimeAndZone(jsonIso8601msTime, iso8601msTime) {
			t.Errorf("Unmarshaled time = %v, want %v", jsonIso8601msTime, iso8601msTime)
		}
	}
}

func equalTimeAndZone(aa, bb Time) bool {
	a, b := time.Time(aa), time.Time(bb)
	aname, aoffset := a.Zone()
	bname, boffset := b.Zone()
	return a.Equal(b) && aoffset == boffset && aname == bname
}
