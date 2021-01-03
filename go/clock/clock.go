package clock

import (
	"fmt"
	"math"
)

// Clock type represents a clock with hours and minutes
type Clock struct {
	hour int
	min  int
}

// rollOver adjusts the minutes and hours as required.
func rollOver(h, m int) (int, int) {
	if m > 59 {
		h += (m / 60)
		m = m % 60
	} else if m < 0 && m > -60 {
		m = m * (-1)
		h--
		m = 60 - m%60
	} else if m == -60 {
		h--
		m = 0
	} else if m < -60 {
		m = m * (-1)
		h -= (int(math.Round(float64(m/60))) + 1)
		m = 60 - m%60
	}

	if h > 23 {
		h = h % 24
	} else if h < 0 {
		h = h * (-1)
		h = 24 - h%24
	}

	if h == 24 {
		h = 0
	}

	return h, m
}

//New initializes a new Clock.
func New(h, m int) Clock {
	h, m = rollOver(h, m)
	return Clock{h, m}
}

//String function returns the string version of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

//Add adds the given minutes and returns the final Clock
func (c Clock) Add(m int) Clock {
	c.hour, c.min = rollOver(c.hour, c.min+m)
	return c
}

//Subtract subtracts the given minutes and returns the final Clock
func (c Clock) Subtract(m int) Clock {
	c.hour, c.min = rollOver(c.hour, c.min-m)
	return c
}
