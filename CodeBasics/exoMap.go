package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := map[string]int{}
	words := strings.Fields(s)
	for _, w := range words {
		c, ok := res[w]
		if ok {
			c++
		} else {
			c = 1
		}
		res[w] = c
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
