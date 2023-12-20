package convertdata

import (
	"time"
)

func DateToString(date time.Time) string {
	return date.String()
}

func CreateDate(year, month, day, hour, min int) time.Time {
	months := map[int]time.Month{
		1:  time.January,
		2:  time.February,
		3:  time.March,
		4:  time.April,
		5:  time.May,
		6:  time.June,
		7:  time.July,
		8:  time.August,
		9:  time.September,
		10: time.October,
		11: time.November,
		12: time.December,
	}
	return time.Date(year, months[month], day, hour, min, 0, 0, time.UTC)
}

func StringToDate(str string) time.Time {
	t, _ := time.Parse(time.DateTime, str)
	return t
}
