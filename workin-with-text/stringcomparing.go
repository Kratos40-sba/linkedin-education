package main

import (
	"fmt"
	"strings"
)

type Letter struct {
	GreekSymbol string
	English string
}
var letters = []Letter{
	{"",""},
}
func englishFor(greek string) (string,error)  {
	for _ , letter := range letters {
		if strings.EqualFold(greek,letter.GreekSymbol) {
			return letter.English , nil
		}
	}
	return "",fmt.Errorf("unknown greek letter : %#v",greek)
}
func main()  {
	// test it with sigma in 3 forms
}