package main

import (
	"fmt"
)

func countMatches(song string, queries [][]int) int {
	count := 0
	for _, query := range queries {
		i, j, k := query[0], query[1], query[2]
		if song[i-1:i-1+k] == song[j-1:j-1+k] {
			count++
		}
	}
	return count
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	var songText string
	fmt.Scan(&songText)

	queries := make([][]int, M)
	for i := 0; i < M; i++ {
		var iVal, jVal, kVal int
		fmt.Scan(&iVal, &jVal, &kVal)
		queries[i] = []int{iVal, jVal, kVal}
	}

	result := countMatches(songText, queries)
	fmt.Println(result)
}
