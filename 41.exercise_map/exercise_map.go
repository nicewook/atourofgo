package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordmap := make(map[string]int)
	for _, v := range strings.Fields(s) {
		wordmap[v] += 1
	}
	return wordmap
}

func main() {
	wc.Test(WordCount) // it can test implementation
}
