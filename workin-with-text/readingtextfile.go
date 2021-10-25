package main

import (
	"bufio"
	"io"
	"strings"
)

func grep(r io.Reader , term string) ([]string,error)  {
	var matches []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		if strings.Contains(s.Text(),term) {
			matches = append(matches,s.Text())
		}
	}
	if err := s.Err() ; err != nil {
		return nil, err
	}
	return matches,nil
}
