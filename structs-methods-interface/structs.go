package main

import (
	"fmt"
	"log"
	"time"
)
type Action string

const  (
	OPEN = Action("open")
	CLOSE = Action("close")
)
type Event struct {
	ID string
	Time time.Time
}
type DoorEvent struct {
	Event
	Action Action// Open , Close
}
type TemperatureEvent struct {
	Event
	Value float64
}

func NewDoorEvent(id string , time time.Time , action Action) (*DoorEvent,error)  {
	if id == ""{
		return nil,fmt.Errorf("empty id")
	}
	event := DoorEvent{
		Event:  Event{id,time},
		Action: action,
	}
	return &event,nil // return pointer to something allocated into the stack
	// escape analysis move from the stack to heap
}
func main() {
	event, err := NewDoorEvent("front-door", time.Now(), OPEN)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n\n", event)
}