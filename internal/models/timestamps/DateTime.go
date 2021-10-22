package timestamps

import (
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

const dateTimeLayout = "02.01.2006 15:04"

func NewDateTime(string_ string) (*DateTime, error) {
	time_, err := time.Parse(dateTimeLayout, string_)
	if err != nil {
		return nil, err
	}

	dateTime := DateTime(time_)
	return &dateTime, nil
}

func (dateTime *DateTime) String() string {
	time_ := time.Time(*dateTime)
	return fmt.Sprintf("%q", time_.Format(dateTimeLayout))
}

func (dateTime *DateTime) UnmarshalJSON(b []byte) (err error) {
	var string_ = strings.Trim(string(b), `"`)
	time_, err := time.Parse(dateTimeLayout, string_)
	*dateTime = DateTime(time_)
	return
}

func (dateTime *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(dateTime.String()), nil
}
