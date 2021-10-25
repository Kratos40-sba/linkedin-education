package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Trade struct {
	Symbol string
	Volume int
	Price float64
}

func genReport(w io.Writer , trades []Trade)  {
	for i , t := range trades {
		log.Printf("%d : %#v",i , t)
		// .. 2 : main.Trade{Symbol: ...
		fmt.Fprintf(w,"%-10s %04d %.2f\n",t.Symbol,t.Volume,t.Price)
		// MSFT 0231 234.75
		time.Sleep(1*time.Second)
	}
}
func main()  {
	log.SetPrefix("LOG : ")
	trades := []Trade{
		{"MSFT",231,234.57},
		{"TSLA",132,686.75},
		{"APPL",1,399100},
	}
	genReport(os.Stdout,trades)

}