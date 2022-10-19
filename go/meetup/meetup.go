package meetup

import (
	"fmt"
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Fifth  WeekSchedule = "fifth"
	Teenth WeekSchedule = "teenth"
	Last   WeekSchedule = "last"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	day := 0
	fmt.Println(wSched, wDay, month, year)
	fmt.Println(time.Now().Year())
	cy := time.Now().Year()
	cm := time.Now().Month()
	cd := time.Now().Day()
	wd := time.Now().Weekday()

	fmt.Println(cy, cm, cd, wd)

	start := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)

	fmt.Println(start)
	return day
}
