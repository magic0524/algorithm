package main

import (
	"fmt"
)

func main() {
	a := []int{5, 5, 3, 2, 1, 9, 3, 5, 7, 1}
	fmt.Println("Rainfall: ", a)

	trapDoublePointer(a)
	trapStack(a)
}

func trapDoublePointer(a []int) int {
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

func trapStack(a []int) int {
	if len(a) < 3 {
		return 0
	}

	stack := make([]int, 0)
	ans := 0

	for i := 0; i < len(a); i++ {
		for len(stack) > 0 && a[i] > a[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			distance := i - stack[len(stack)-1] - 1
			boundedHeight := min(a[i], a[stack[len(stack)-1]]) - a[top]
			ans += distance * boundedHeight
		}

		stack = append(stack, i)
	}

	fmt.Printf("ans: %d\n", ans)

	return ans
}
