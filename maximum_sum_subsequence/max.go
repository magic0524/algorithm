package max

import (
	"fmt"
	"math"
)

func findLargestSumSubarray(arr []int, k int) int {
	maxSum := math.MinInt32
	dp := make([]int, len(arr))
	copy(dp, arr)

	for i := 0; i < len(arr); i++ {
		for j := max((i - k), 0); j < i; j++ {
			dp[i] = max(dp[i], dp[j]+arr[i])
		}
		maxSum = max(maxSum, dp[i])
	}

	fmt.Printf("max sum: %d\n", maxSum)

	return maxSum
}
