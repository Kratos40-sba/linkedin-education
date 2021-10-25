package main

import (
	"fmt"
	"time"
)

func main()  {
	// if u dont care about time zones use utc
	chicagoLocation , err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Printf("error : %v",err)
		return
	}
	chicagoTime := time.Date(2021,time.October,21,19,60,0,0,chicagoLocation)
	fmt.Println("Chicago : ",chicagoTime)
	nyc , err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("error : %s ",err)
		return
	}
	nycTime := chicagoTime.In(nyc)
	fmt.Println("NYC : ",nycTime)
}
