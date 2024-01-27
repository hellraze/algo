package main

import "fmt"

func main() {
	var (
		n, k int
		arr  []int
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		arr = append(arr, k)
	}
	print_preorder(arr, 0, len(arr)-1)
}

func print_preorder(arr []int, start int, end int) {
	if start <= end {
		half := start + (end-start)/2
		fmt.Print(arr[half], " ")
		print_preorder(arr, start, half-1)
		print_preorder(arr, (half + 1), end)
	}
}
