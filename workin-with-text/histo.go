package main

import (
	"bufio"
	"os"
	"regexp"
)
/*
example : 15247842780;go test
...
 */
var cmdRe = regexp.MustCompile(`;go ([a-z]+)`)
// cmdFrea returns the frequency of "go" subcommand usage in zsh history
func cmdFrea(fileName string) (map[string]int,error)  {
	file , err := os.Open(fileName)
	if err != nil {
		return nil , err
	}
	defer file.Close()
	freqs := make(map[string]int)
	s:= bufio.NewScanner(file)
	for s.Scan() {
		matches := cmdRe.FindStringSubmatch(s.Text())
		if len(matches) == 0 {
			continue
		}
		cmd := matches[1]
		freqs[cmd]++

	}
	if err := s.Err() ; err != nil {
		return nil, err
	}

	return freqs,nil
}
