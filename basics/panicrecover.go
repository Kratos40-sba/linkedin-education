package main

import "fmt"

func last2Values(num []int)[]int  {
	res := make([]int,0)
	res = append(res,num[len(num)-2:]...)
	return res
}
func safeLast2Values(num []int) (res []int,err error)  {
	defer func() {
		if e := recover();e != nil {
			err = fmt.Errorf("%v",e)
		}
	}()
	res = last2Values(num)
	return res,nil
}
func main()  {
	l := []int{1}
	fmt.Println(safeLast2Values(l))
}

