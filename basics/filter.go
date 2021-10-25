package main

import "fmt"

func filter(pred func(int)bool,values []int) []int  {
	res := make([]int , 0 )
	for _,v := range values {
		if pred(v) {
			res = append(res,v)
		}
	}
	return res
}
func isOdd(n int) bool  {
	return n%2==1
}
func main()  {
	values := []int{1,2,3,5,9,8,10}
	fmt.Println(filter(isOdd,values)) // 1,3,5,9
}