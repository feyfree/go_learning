package struct_test

import (
	"fmt"
	"testing"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func NewDate(year int, month int, day int) *Date {
	return &Date{Year: year, Month: month, Day: day}
}

func (d *Date) SetYear(y int) {
	d.Year = y
}

func (d *Date) SetMonth(m int) {
	d.Month = m
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func TestDate(t *testing.T) {
	date := NewDate(2020, 9, 20)
	fmt.Println(*date)
	date.Day = 21
	fmt.Println(*date)
}
