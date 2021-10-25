package main

import (
	"fmt"
	"unicode/utf8"
)

func lineLength(words []string) int  {
	total := 0
	for _ , word := range words {
		total += utf8.RuneCountInString(word)
		//total += len(word)
	}
	numSpaces := len(words) - 1
	return total + numSpaces
}
func main()  {
	words := []string{"<","Don't","Panic",">"} // 15 with space
	fmt.Println(lineLength(words))
}
