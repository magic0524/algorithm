package goodnum

import "fmt"

func findGood(arr []int, k int) (result bool) {
	defer func() {
		fmt.Printf("arr: %v, k: %d, result: %v\n", arr, k, result)
	}()

	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}

	for i := 0; i < len(arr); i++ {
		dp[i][i] = arr[i]
	}

	for l := 2; l < len(arr); l++ {
		for i := 0; i < len(arr)-l+1; i++ {
			j := i + l - 1
			x := dp[i][j-1] + arr[j]
			if x%k == 0 {
				return true
			}

			dp[i][j] = x
		}
	}

	return false
}
