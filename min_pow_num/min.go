package numsum

import (
	"fmt"
)

// dp公式：dp[i] = min(dp[i-j]+dp[j]) 1<=j<=i/2
func min(a int) int {
	dp := make([]int, a+1)
	dp[0] = 0

	for i := 1; i <= a; i++ {
		min := i
		for j := 1; j < i/2+1; j++ {
			if j*j == i {
				min = 1
				break
			}
			if min > dp[i-j]+dp[j] {
				min = dp[i-j] + dp[j]
			}
		}
		dp[i] = min
	}
	fmt.Printf("a = %d, min = %v\n", a, dp[a])
	return dp[a]
}
