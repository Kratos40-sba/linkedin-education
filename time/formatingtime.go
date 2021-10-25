package main

import (
	"fmt"
	"time"
)

func main()  {
	lennon := time.Date(1940,time.August,9,18,20,15,0,time.UTC)
	fmt.Println(lennon.Format("2006-01-02"))
	fmt.Println(lennon.Format("Mon, Jan 02"))
	fmt.Println(lennon.Format(time.RFC3339Nano))
}
