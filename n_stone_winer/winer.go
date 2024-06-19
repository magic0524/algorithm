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

// 只要数量是k的倍数，那先手就输，因为无论先手拿多少个，后手都可以拿k个，使得剩下的数量是k的倍数
func winer_n(n int, k int) bool {
	return n%k != 0
}
