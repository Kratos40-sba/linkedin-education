package main

import "fmt"

func freq(word string) map[string]int  {
	freqs := make(map[string]int)
	for _,w := range word {
		freqs[string(w)]++
	}
	return freqs
}
var myString = "aaaabccdef"

func searchInString(m map[string]int, s string) bool   {
	// comma ok pattern in maps
	_ , ok := m[s] ; if ok {
		return ok
	}
	return false
}
func main()  {
	fmt.Println(searchInString(freq(myString),"z"))
}