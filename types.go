package rawgSdkGo

import (
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

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
