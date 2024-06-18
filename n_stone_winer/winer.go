package winer

import "fmt"

func winer_3(n int) int {
	fmt.Printf("n = %d\n", n)
	dp := make([]int, n+1)
	dp[0] = -1
	dp[1] = 1
	dp[2] = 1
	dp[3] = 1

	for i := 4; i <= n; i++ {
		if min(dp[i-1], dp[i-2], dp[i-3]) == 1 {
			dp[i] = 0
		} else {
			dp[i] = 1
		}
	}

	fmt.Printf("dp[n] = %d\n", dp[n])
	return dp[n]
}

func winer_n(n int, k int) bool {
	return n%k != 0
}
