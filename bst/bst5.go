package main

import (
	"fmt"
)

func countScores(n int, solutions [][]string) []int {
	scores := []int{0, 0, 0}
	counts := make(map[string]int)

	for _, solution := range solutions {
		seen := make(map[string]struct{})
		for _, file := range solution {
			if _, ok := seen[file]; !ok {
				counts[file]++
				seen[file] = struct{}{}
			}
		}
	}

	for i, solution := range solutions {
		for _, file := range solution {
			if counts[file] == 1 {
				scores[i] += 3
			} else if counts[file] == 2 {
				scores[i]++
			}
		}
	}

	return scores
}

func main() {
	var n int
	fmt.Scan(&n)

	studentsSolutions := make([][]string, 3)
	for i := 0; i < 3; i++ {
		solutions := make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&solutions[j])
		}
		studentsSolutions[i] = solutions
	}

	finalScores := countScores(n, studentsSolutions)
	fmt.Println(finalScores[0], finalScores[1], finalScores[2])
}
