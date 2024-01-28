package main

import (
	"fmt"
)

func binSearch(n int, l, r float64) float64 {
	x := -1.0
	h := r

	for (r - l) > 1e-6 {
		mid := (l + r) / 2
		h0, hi := h, mid
		greater := hi > 0

		for i := 3; i <= n; i++ {
			hn := 2*hi - h0 + 2
			greater = (hn > 0) && greater
			h0, hi = hi, hn
		}

		if greater {
			r = mid
			x = hi
		} else {
			l = mid
		}
	}

	return x
}

func main() {
	var n int
	var A float64

	fmt.Scan(&n, &A)

	result := binSearch(n, 0, A)

	fmt.Printf("%.2f\n", result)
}
