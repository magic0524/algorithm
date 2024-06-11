package main

import (
	"fmt"
)

func main() {
	a := []int{5, 5, 3, 2, 1, 9, 3, 5, 7, 1}
	fmt.Println("Rainfall: ", a)

	rain(a)
}

func rain(a []int) int {
	if len(a) < 3 {
		return 0
	}

	l := 0
	r := len(a) - 1

	maxLeft := a[l]
	maxRight := a[r]

	ans := 0

	for l < r {
		if maxLeft < maxRight {
			ans += maxLeft - a[l]
			l++
			if a[l] > maxLeft {
				maxLeft = a[l]
			}
		} else {
			ans += maxRight - a[r]
			r--
			if a[r] > maxRight {
				maxRight = a[r]
			}
		}
	}

	fmt.Printf("ans: %d\n", ans)

	return ans
}
