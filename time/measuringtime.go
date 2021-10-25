package main

import (
	"fmt"
	"log"
	"time"
)

func timeIt(name string) func()  {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		log.Printf("%s took %v \n", name , duration)
	}
}
func dot(v1,v2 []float64) (float64,error)  {
	defer timeIt("dot")()
	if len(v1) != len(v2) {
		return 0 , fmt.Errorf("dot of different size vectors")
	}
	d := 0.0
	// sleeping
	//time.Sleep(1*time.Second)
	for i,val1 := range v1 {
		val2 := v2[i]
		d += val2 * val1
	}
	return d , nil
}
func main()  {
	vector1 := []float64{2,5,8,9,7,4,1,5,6,9,7,14}
	vector2 := []float64{214,55,1.8,19,7,4,1,5,6,9,7,1400}
	fmt.Println(dot(vector1, vector2))
}