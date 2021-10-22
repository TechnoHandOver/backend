package timestamps

import (
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

const fullDateLayout = "02.01.2006 15:04"

func (dateTime *DateTime) String() string {
	time_ := time.Time(*dateTime)
	return fmt.Sprintf("%q", time_.Format(fullDateLayout))
}

func (dateTime *DateTime) UnmarshalJSON(b []byte) (err error) {
	var string_ = strings.Trim(string(b), `"`)
	time_, err := time.Parse(fullDateLayout, string_)
	*dateTime = DateTime(time_)
	return
}

func (dateTime *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(dateTime.String()), nil
}
