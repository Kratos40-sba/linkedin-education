package main

import (
	"fmt"
	"time"
)

func convertTimeZone(ts , from , to string) (string,error)  {
	fromTZ ,err := time.LoadLocation(from)
	if err != nil {
		fmt.Println("fromTZ")
		return "",err
	}
	toTZ ,err := time.LoadLocation(to)
	if err != nil {
		fmt.Println("toTZ")
		return "",err
	}
	const format = "2006-01-02T15:04"
	fromTime , err := time.ParseInLocation(format,ts,fromTZ)
	if err != nil {
		fmt.Println("fromTime")
		return "",err
	}
	toTime := fromTime.In(toTZ)
	return toTime.Format(format), nil

}
func main()  {
	ts := "2021-03-08T19:12"
	out , err := convertTimeZone(ts , "America/Chicago","Asia/Jerusalem")
	if err != nil {
		fmt.Printf("error : %s \n",err)
		return
	}
	fmt.Println(out)
}
