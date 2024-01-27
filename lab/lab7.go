package main

import (
	"fmt"
)

func findPrefixLen(spell string) []int {
	n := len(spell)
	prefixLen := make([]int, n)

	for i := 1; i < n; i++ {
		j := prefixLen[i-1]
		for j > 0 && spell[i] != spell[j] {
			j = prefixLen[j-1]
		}
		if spell[i] == spell[j] {
			j++
			prefixLen[i] = j
		}
	}

	return prefixLen
}

func main() {
	var spell string
	fmt.Scan(&spell)
	prefixLen := findPrefixLen(spell)

	for i := 0; i < len(prefixLen); i++ {
		fmt.Print(prefixLen[i], " ")
	}
}
