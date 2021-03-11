package iso8601ms

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Run("UTC format length should be 26", func(t *testing.T) {
		iso8601msTime := Time(time.Now().UTC())
		jsonBytes, err := json.Marshal(iso8601msTime)
		if err != nil {
			t.Errorf("error = %v, want nil", err)
		}
		if len(jsonBytes) != 26 {
			t.Errorf("invalid json string length %v, want %v", len(jsonBytes), 26)
		}
	})

	t.Run("Local format length should be 31", func(t *testing.T) {
		iso8601msTime := time.Date(1996, 12, 19, 16, 39, 57, 1*1e6, time.FixedZone("UTC-7", -7*60*60))
		jsonBytes, err := json.Marshal(iso8601msTime)
		if err != nil {
			t.Errorf("error = %v, want nil", err)
		}
		if len(jsonBytes) != 31 {
			t.Errorf("invalid json string length %v, want %v", len(jsonBytes), 31)
		}
	})
}

func TestUnmarshal(t *testing.T) {
	t.Run("valid cases", func(t *testing.T) {
		var jsonIso8601msTime Time
		var err error

		err = json.Unmarshal([]byte(`"2021-03-11T01:29:34.317-07:00"`), &jsonIso8601msTime)
		if err != nil {
			t.Errorf("json.Unmarshal error = %v, want nil", err)
		}

		err = json.Unmarshal([]byte(`"2021-03-11T01:29:34.317Z"`), &jsonIso8601msTime)
		if err != nil {
			t.Errorf("json.Unmarshal error = %v, want nil", err)
		}

		err = json.Unmarshal([]byte(`"2021-03-11T01:29:34Z"`), &jsonIso8601msTime)
		if err != nil {
			t.Errorf("json.Unmarshal error = %v, want nil", err)
		}

		err = json.Unmarshal([]byte(`"2021-03-11T01:29:34.999999Z"`), &jsonIso8601msTime)
		if err != nil {
			t.Errorf("json.Unmarshal error = %v, want nil", err)
		}

		err = json.Unmarshal([]byte(`"2021-03-11T01:29:34.999999999Z"`), &jsonIso8601msTime)
		if err != nil {
			t.Errorf("json.Unmarshal error = %v, want nil", err)
		}

	})

	t.Run("invalid cases", func(t *testing.T) {
		var jsonIso8601msTime Time
		var err error

		err = json.Unmarshal([]byte(`"2021-03-11T01:29:34.317"`), &jsonIso8601msTime)
		if err == nil {
			t.Errorf("got nil error, want parsing error from json.Unmarshal")
		}

		err = json.Unmarshal([]byte(`"2021-03-11T01:29:34"`), &jsonIso8601msTime)
		if err == nil {
			t.Errorf("got nil error, want parsing error from json.Unmarshal")
		}
	})
}

var jsonTests = []struct {
	time time.Time
	json string
}{
	{time.Date(9999, 4, 12, 23, 20, 50, 520*1e6, time.UTC), `"9999-04-12T23:20:50.520Z"`},
	{time.Date(1996, 12, 19, 16, 39, 57, 0, time.FixedZone("UTC-4", -4*60*60)), `"1996-12-19T16:39:57.000-04:00"`},
	{time.Date(1996, 12, 19, 16, 39, 57, 1*1e6, time.FixedZone("UTC-7", -7*60*60)), `"1996-12-19T16:39:57.001-07:00"`},
	{time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), `"0000-01-01T00:00:00.000Z"`},
}

func TestIso8601msJSON(t *testing.T) {
	for _, tt := range jsonTests {
		iso8601msTime := Time(tt.time)
		var jsonIso8601msTime Time

		if jsonBytes, err := json.Marshal(iso8601msTime); err != nil {
			t.Errorf("%v json.Marshal error = %v, want nil", iso8601msTime, err)
		} else if string(jsonBytes) != tt.json {
			t.Errorf("%v JSON = %#q, want %#q", iso8601msTime, string(jsonBytes), tt.json)
		} else if err = json.Unmarshal(jsonBytes, &jsonIso8601msTime); err != nil {
			t.Errorf("%v json.Unmarshal error = %v, want nil", iso8601msTime, err)
		} else if !equalTime(jsonIso8601msTime, iso8601msTime) {
			t.Errorf("Unmarshaled time = %v, want %v", jsonIso8601msTime, iso8601msTime)
		}
	}
}

func equalTime(aa, bb Time) bool {
	a, b := time.Time(aa), time.Time(bb)
	return a.Equal(b)
}
