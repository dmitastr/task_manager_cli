package datetimeformattter

import (
	"fmt"
	"strings"
	"time"
)

const (
	dttmFormat string = "2006-01-02 15:04:05"
)

type DateTime time.Time


func DateTimeNow() DateTime {
	return DateTime(time.Now())
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	// layout := fmt.Sprintf(`%s`, dttmFormat)
	date, err := time.Parse(dttmFormat, s)
	if err != nil {
		return err
	}
	*t = DateTime(date)
	return 
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	tm := time.Time(t)
	formattedTime := fmt.Sprintf("\"%s\"", tm.Format(dttmFormat))
	return []byte(formattedTime), nil
}
