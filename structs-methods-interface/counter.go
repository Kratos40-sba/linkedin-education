package main

import (
	"fmt"
	"log"
)

type (
	ClickEvt struct {

	}
	HoverEvt struct {

	}
)
var eventsCounts = make(map[string]int) // type -> count
func recordEvents(evt interface{})  {
	switch evt.(type) {
	case *ClickEvt:
		eventsCounts["click"]++
	case *HoverEvt:
		eventsCounts["hover"]++
	default:
		log.Printf("warning : unkown event : %#v of type %T\n",evt,evt)
	}
}
func main()  {
	recordEvents(&ClickEvt{})
	recordEvents(&HoverEvt{})
	recordEvents(&ClickEvt{})
	recordEvents(335)
	fmt.Println("Events Counts : ",eventsCounts)

}
