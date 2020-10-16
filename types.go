package rawg

import (
	"strings"
	"time"
)

// DateTime: special type to unmarshal dates from any formats
type DateTime struct {
	time.Time
}

// UnmarshalJSON converts json to time.Time
func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return nil
	}
	format := "2006-01-02T15:04:05"
	if strings.HasSuffix(s, "Z") {
		format = "2006-01-02T15:04:05Z"
	}

	if len(s) == 10 {
		format = "2006-01-02"
	}

	t, err := time.Parse(format, s)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}
