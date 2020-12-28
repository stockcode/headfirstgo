package main

import (
	"fmt"
	"headfirstgo/calendar"
)

func main()  {
	date := calendar.Date{}
	date.SetYear(2019)
	date.SetMonth(5)
	date.SetDay(27)
	fmt.Println(date.Year())
}
