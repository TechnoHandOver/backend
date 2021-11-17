package timestamps

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Time time.Time

const timeLayout = "15:04"

func NewTime(string_ string) (*Time, error) {
	time_, err := time.Parse(timeLayout, string_)
	if err != nil {
		return nil, err
	}

	time__ := Time(time_)
	return &time__, nil
}

func (time_ *Time) String() string {
	time__ := time.Time(*time_)
	return fmt.Sprintf("%q", time__.Format(timeLayout))
}

func (time_ *Time) UnmarshalJSON(b []byte) (err error) {
	var string_ = strings.Trim(string(b), `"`)
	time__, err := time.Parse(timeLayout, string_)
	*time_ = Time(time__)
	return
}

func (time_ *Time) MarshalJSON() ([]byte, error) {
	return []byte(time_.String()), nil
}

func ValidateTime(field reflect.Value) interface{} {
	if time_, ok := field.Interface().(Time); ok {
		if _, err := NewTime(strings.Trim(time_.String(), `"`)); err == nil {
			return true
		}
	}
	return false //TODO: точно true и false эта функция должна возвращать?; не только тут такой валидатор
}
