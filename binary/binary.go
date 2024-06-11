package main

import "fmt"

func main() {
	a := []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	fmt.Printf("a: %v\n", a)

	fmt.Printf("Index of 1: %d\n", binarySearch(a, 1))
	fmt.Printf("Index of 9: %d\n", binarySearch(a, 9))
	fmt.Printf("Index of -1: %d\n", binarySearch(a, -1))
	fmt.Printf("Index of 10: %d\n", binarySearch(a, 10))
	fmt.Printf("Index of 0: %d\n", binarySearch(a, 0))
}

func binarySearch(a []int, k int) int {
	fmt.Printf("k = %d\n", k)
	mid := len(a) / 2
	left := 0
	right := len(a) - 1

	for left < right {
		if k == a[mid] {
			return mid
		}

		if left == mid || right == mid {
			break
		}

		if a[mid] < a[right] {
			if k > a[mid] && k <= a[right] {
				left = mid
				mid = (mid + right) / 2
			} else {
				right = mid
				mid = (mid + left) / 2
			}
		} else {
			if k < a[mid] && k >= a[left] {
				right = mid
				mid = (mid + left) / 2
			} else {
				left = mid
				mid = (mid + right) / 2
			}
		}
	}

	return -1
}
