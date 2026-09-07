package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
    h int
    m int
}

func Format(h, m int) Clock {
    m += h * 60
    h = 0

    if m < 0 {
        m = -m
        for m >= 60 {
            h -= 1
            m -= 60
        }
        h -= 1
        m = 60 - m
    }

    h += m / 60
    m = m % 60

    if h < 0 {
        h = 24 + h % 24
    }

    h = h % 24
    return Clock {h, m}
}

func New(h, m int) Clock {
    formClock := Format(h, m)
	return formClock
}

func (c Clock) Add(m int) Clock {
    formClock := Format(c.h, c.m + m)
	return formClock
}

func (c Clock) Subtract(m int) Clock {
    formClock := Format(c.h, c.m - m)
	return formClock
}

func (c Clock) String() string {
    formClock := Format(c.h, c.m)
    
    return fmt.Sprintf("%02d:%02d", formClock.h, formClock.m)
}
