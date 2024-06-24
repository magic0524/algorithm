package binary

import "fmt"

func binarySearch(a []int, k int) (res int) {
	defer func() {
		fmt.Printf("binary search a = %v\tk = %d\tres = %d\n", a, k, res)
	}()

	if len(a) == 1 {
		if a[0] == k {
			return 0
		}
		return -1
	}

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

func binarySearchMin(a []int) (res int) {
	defer func() {
		fmt.Printf("binary search min a = %v\tres = %d\n", a, res)
	}()

	if len(a) <= 0 {
		return -1
	}

	if len(a) == 1 {
		return a[0]
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			return a[1]
		}
		return a[0]
	}

	mid := len(a) / 2
	start := 0
	end := len(a)

	for start < end {
		if a[mid] > a[start] {
			// 左边是递增
			start = mid
			mid = (end-start)/2 + start
		} else {
			// 左边是递增
			end = mid
			mid = (end-start)/2 + start
		}
	}

	return a[mid+1]
}
