package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int) // this would be a returned answer
	str := strings.Fields(s)
	
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
