package main

import (
	"fmt"
	"time"
)

func isBusinessDay(date time.Time) bool  {
	wday := date.Weekday()
	if wday == time.Saturday || wday == time.Friday {
		return false
	}
	return true
}
func nextBusinessDay(date time.Time) time.Time  {
	const day = 24*time.Hour
	for {
		date = date.Add(day)
		if isBusinessDay(date){
			break
		}
	}
	return date
}
func main()  {
	date := time.Date(2021,time.December,31,0,0,0,0,time.UTC)
	fmt.Println(date,date.Weekday()) // friday
	nbd := nextBusinessDay(date)
	fmt.Println(nbd,nbd.Weekday()) // sunday

}