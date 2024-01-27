package main

import (
	"fmt"
)

func left(arr []int, target int) int {
	start, end := 0, len(arr)-1
	result := -1
	for start <= end {
		mid := (start + end) / 2
		midValue := arr[mid]
		if midValue == target {
			result = mid
			end = mid - 1
		} else if midValue < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

func right(arr []int, target int) int {
	start, end := 0, len(arr)-1
	result := -1
	for start <= end {
		mid := (start + end) / 2
		midValue := arr[mid]
		if midValue == target {
			result = mid
			start = mid + 1
		} else if midValue < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

func main() {
	var n, m int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&m)

	queries := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&queries[i])
	}

	for _, target := range queries {
		first := left(arr, target)
		last := right(arr, target)
		if first != -1 && last != -1 {
			fmt.Println(first+1, last+1)
		} else {
			fmt.Println(-1, -1)
		}
	}
}
