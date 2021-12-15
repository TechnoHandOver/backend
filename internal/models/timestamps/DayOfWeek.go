package timestamps

import (
	"errors"
	"fmt"
)

type DayOfWeek string

const (
	DayOfWeekMonday    DayOfWeek = "Mon"
	DayOfWeekTuesday   DayOfWeek = "Tue"
	DayOfWeekWednesday DayOfWeek = "Wed"
	DayOfWeekThursday  DayOfWeek = "Thu"
	DayOfWeekFriday    DayOfWeek = "Fri"
	DayOfWeekSaturday  DayOfWeek = "Sat"
	DayOfWeekSunday    DayOfWeek = "Sun"
)

func (dayOfWeek *DayOfWeek) ToUint32() (uint32, error) {
	switch *dayOfWeek {
	case DayOfWeekMonday:
		return 1, nil
	case DayOfWeekTuesday:
		return 2, nil
	case DayOfWeekWednesday:
		return 3, nil
	case DayOfWeekThursday:
		return 4, nil
	case DayOfWeekFriday:
		return 5, nil
	case DayOfWeekSaturday:
		return 6, nil
	case DayOfWeekSunday:
		return 7, nil
	default:
		return 0, errors.New(fmt.Sprintf("Cannot parse day of week %s\n", *dayOfWeek))
	}
}
