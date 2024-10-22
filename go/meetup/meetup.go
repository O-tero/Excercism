package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day function calculates the correct day based on the week schedule, weekday, month, and year
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var start, end int

	switch wSched {
	case First:
		start = 1
		end = 7
	case Second:
		start = 8
		end = 14
	case Third:
		start = 15
		end = 21
	case Fourth:
		start = 22
		end = 28
	case Last:
		t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC) 
		start = t.Day() - 6
		end = t.Day()
	case Teenth:
		start = 13
		end = 19
	}

	
	for day := start; day <= end; day++ {
		t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if t.Weekday() == wDay {
			return day
		}
	}
	return -1
}
