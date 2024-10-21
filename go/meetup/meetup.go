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

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
    switch wSched {
    case First:
        return firstWeekday(month, year, wDay)
    case Second:
        return firstWeekday(month, year, wDay) + 7
    case Third:
        return firstWeekday(month, year, wDay) + 14
    case Fourth:
        return firstWeekday(month, year, wDay) + 21
    case Last:
        return findLastDay(month, year, wDay)
    case Teenth:
        return findTeenthDay(month, year, wDay)
    default:
        return -1
    }
}

func firstWeekday(month time.Month, year int, wDay time.Weekday) int {
    t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
    for t.Weekday() != wDay {
        t = t.AddDate(0, 0, 1) 
    }
    return t.Day()
}

func findTeenthDay(month time.Month, year int, wDay time.Weekday) int {
    for day := 13; day <= 19; day++ {
        t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
        if t.Weekday() == wDay {
            return day
        }
    }
    return -1 
}

func findLastDay(month time.Month, year int, wDay time.Weekday) int {
    t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
    for t.Weekday() != wDay {
        t = t.AddDate(0, 0, -1) 
    }
    return t.Day()
}
