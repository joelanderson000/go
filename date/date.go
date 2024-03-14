// Basic date object that stores day and time as integer values
package date

type Month int

const (
	JAN Month = iota
	FEB
	MAR
	APR
	MAY
	JUN
	JUL
	AUG
	SEP
	OCT
	NOV
	DEC
)

type DayOfYear struct {
	year  int
	month Month
	day   int
}

type Time struct {
	hour int
	min  int
	sec  int
}

type Date struct {
	day  DayOfYear
	time Time
}

func NewDayOfYear(year int, month Month, day int) *DayOfYear {
	return &DayOfYear{year, month, day}
}

func NewTime(hour int, min int, sec int) *Time {
	return &Time{hour, min, sec}
}

func NewDefaultTime() *Time {
	return &Time{}
}

func NewDate(day DayOfYear) *Date {
	return &Date{day, *NewDefaultTime()}
}

func NewDateWithTime(day DayOfYear, time Time) *Date {
	date := new(Date)
	date.day, date.time = day, time

	return date
}
