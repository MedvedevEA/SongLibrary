package helpers

import (
	"fmt"
	"strings"
	"time"
)

type DateDDMMYYYY struct {
	time.Time
}

const layout = "02-01-2006"

func (c *DateDDMMYYYY) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

func (c DateDDMMYYYY) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}
