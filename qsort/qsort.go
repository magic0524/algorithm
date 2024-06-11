package main

import (
	"fmt"
)

func main() {
	a := []int{4, 3, 6, 1, 8, 3, 5}
	fmt.Printf("Before qsort: %v\n", a)

	fmt.Printf("After qsort: %v\n", qsort(a))
}

func qsort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	left := 0
	right := len(a) - 1
	pivot := a[0]

	for left < right {
		for left < right && a[right] >= pivot {
			right--
		}
		a[left] = a[right]

		for left < right && a[left] <= pivot {
			left++
		}
		a[right] = a[left]
	}

	a[left] = pivot

	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
