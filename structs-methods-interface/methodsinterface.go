package main

import "fmt"

type (
	Thermostat struct {
		id string
		value float64
	}
	Camera struct {
		id string
	}
	Sensor interface {
		ID() string
		Kind() string
	}
)

func(c *Camera) ID () string  {
	return c.id
}
func(*Camera) Kind () string {
	return "Camera"
}
func(t *Thermostat) ID () string  {
	return t.id
}
func(*Thermostat) Kind() string  {
	return "Thermostat"
}
func PrintAllSensors(sensors []Sensor)  {
	for _,s := range sensors {
		fmt.Printf("%s <%s> \n",s.ID(),s.Kind())
	}
}
func main ()  {
	t := Thermostat{
		id:    "Living Room",
		value: 18,
	}
	c := Camera{id: "Baby Room"}
	sensors := []Sensor{&t,&c}
	PrintAllSensors(sensors)
}
