package main

import (
	"fmt"
)

func weekDay(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Wrong day"
	}
}

func weekDay2(number int) string {
	var day string

	switch number {
	case 1:
		day = "Sunday"
	case 2:
		day = "Monday"
	case 3:
		day = "Tuesday"
	case 4:
		day = "Wednesday"
	case 5:
		day = "Thursday"
	case 6:
		day = "Friday"
	case 7:
		day = "Saturday"
	default:
		day = "Wrong day"
	}

	return day
}

func main() {
	fmt.Println("Switch")

	day := weekDay(3)
	fmt.Println(day)

	day2 := weekDay(10)
	fmt.Println(day2)
}
