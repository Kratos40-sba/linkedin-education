package main

import (
	"fmt"
	"time"
)

func main()  {
	template := "June 18, 1942"
	parsedTime , err := time.Parse("January 02, 2006",template)
	if err != nil {
		fmt.Printf("error : %v",err)
	}else{
		fmt.Println(parsedTime) // 1942-06-18 00:00:00 +0000 UTC

	}
	ds := "2700ms"
	d , err := time.ParseDuration(ds)
	if err != nil {
		fmt.Printf("error : %v",err)
	}else{
		fmt.Println(d) // 2.7s
	}
}