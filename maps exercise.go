package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	var fields []string = strings.Fields(s)
	
	m := make(map[string]int)
	
	for i := 0; i < len(fields); i++ {
		elem, ok := m[fields[i]]
		if ok {
			m[fields[i]] = elem + 1
		} else {
			m[fields[i]] = 1
		}
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}