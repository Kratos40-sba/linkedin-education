package main

import (
	"fmt"
	"sort"
)

func main()  {
	var l = []int{1,2,3,5,9}

	fmt.Println(mean(l)) // 4
	fmt.Println(median(toFloats(l)))
}
func mean(nums []int) float64 {
	sum := sum(nums)
	return float64(sum)/float64(len(nums))
}
func sum(nums []int) int  {
	var sum = 0
	for _,n := range nums {
		sum +=n
	}
	return sum
}
/*
 int8 , int16 , int64 -> uint8 , .., uint64
 int , uint -> selon l'architecture de ordinature
 complex64 , complex128
 float64 , float32 => same
 */
// slices
func median(nums []float64) float64  {
	// using copy to not change the input
	vals := make([]float64,len(nums))
	copy(vals , nums)
	sort.Float64s(vals)
	i := len(vals)/2
	// 1 2 3 4 5
	if len(vals)%2 == 1 {
		return vals[i]
	}
	return (vals[i-1]+vals[i])/2


}
func toFloats (ints []int) []float64  {
	floats := make([]float64,len(ints))
	for i,v := range ints {
		floats[i] = float64(v)
	}
	return floats
}