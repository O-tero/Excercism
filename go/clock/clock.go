package clock

import "fmt"

// Clock represents a time without dates.
type Clock struct {
    minutes int
}

func New(h, m int) Clock {
    totalMinutes := h*60 + m
    if totalMinutes < 0 {
        totalMinutes = 1440 + (totalMinutes % 1440)
    }
    totalMinutes %= 1440
    return Clock{minutes: totalMinutes}
}

func (c Clock) Add(m int) Clock {
    totalMinutes := c.minutes + m
    if totalMinutes < 0 {
        totalMinutes = 1440 + (totalMinutes % 1440)
    }
    totalMinutes %= 1440
    return Clock{minutes: totalMinutes}
}

func (c Clock) Subtract(m int) Clock {
    totalMinutes := c.minutes - m
    if totalMinutes < 0 {
        totalMinutes = 1440 + (totalMinutes % 1440)
    }
    totalMinutes %= 1440
    return Clock{minutes: totalMinutes}
}

func (c Clock) String() string {
    hours := c.minutes / 60
    minutes := c.minutes % 60
    return fmt.Sprintf("%02d:%02d", hours, minutes)
}
