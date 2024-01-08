package main

import (
	"fmt"
	"time"
)

// FirstMonday returns the day of the first Monday in the given month.
func FirstMonday(year int, month time.Month) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(int(t.Weekday()))
	return (8-int(t.Weekday()))%7 + 1
}

func main() {
	// for m := 1; m <= 12; m++ {
	// 	fmt.Printf("month : %d , firstMonday day : %d \n", m, FirstMonday(2023, time.Month(m)))
	// 	fmt.Println("==================================")
	// }
	findWeekday(1)
	//
	findWeekday(2)
	findWeekday(3)
	findWeekday(4)
	findWeekday(5)
}

func findWeekday(weekday int) {
	t := time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC)

	diff := ((7 + weekday) - int(t.Weekday())) % 7
	next := t.AddDate(0, 0, diff)
	//fmt.Println(int(t.Weekday()))
	//fmt.Println(weekday)
	fmt.Println(next.Add(time.Duration(36000) * time.Second))
	fmt.Println("==================================")
}
