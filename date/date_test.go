package date

import (
	"testing"
)

// Test year month and day are stored as integers correctly
func TestNewDayOfYear(t *testing.T) {
	year := 1990
	month := AUG
	day := 15

	date := NewDayOfYear(year, month, day)

	if date.year != year {
		t.Errorf("got %q, wanted %q", date.year, year)
	}
	if date.month != month {
		t.Errorf("got %q, wanted %q", date.month, month)
	}
	if date.day != day {
		t.Errorf("got %q, wanted %q", date.day, day)
	}
}

// Test that time is stored as integers correctly
func TestNewTime(t *testing.T) {
	hour := 12
	min := 15
	sec := 45

	time := NewTime(hour, min, sec)

	if time.hour != hour {
		t.Errorf("got %q, wanted %q", time.hour, hour)
	}
	if time.min != min {
		t.Errorf("got %q, wanted %q", time.min, min)
	}
	if time.sec != sec {
		t.Errorf("got %q, wanted %q", time.sec, sec)
	}
}

// Test a default case to initialize and empty time when not specified
func TestNewDate(t *testing.T) {

	date := NewDate(*NewDayOfYear(1985, FEB, 13))
	hour := 0

	if date.time.hour != hour {
		t.Errorf("got %q, wanted %q", date.time.hour, hour)
	}
}

func TestNewDateWithTime(t *testing.T) {

	date := NewDateWithTime(*NewDayOfYear(1985, FEB, 13), *NewTime(12, 0, 30))
	hour := 12

	if date.time.hour != hour {
		t.Errorf("got %q, wanted %q", date.time.hour, hour)
	}
}
