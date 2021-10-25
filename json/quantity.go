package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Quantity struct {
	Value float64
	Unit string
}

func(q *Quantity) MarshalJSON () ([]byte,error) {
	if q.Unit =="" {
		return nil,fmt.Errorf("empty unit")
	}
	text := fmt.Sprintf("%.2f%s",q.Value,q.Unit)
	return json.Marshal(text)
}
func main() {
	q := Quantity{
		Value: 1.78,
		Unit:  "meter",
	}
	json.NewEncoder(os.Stdout).Encode(&q)
}