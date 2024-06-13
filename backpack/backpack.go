package backpack

import (
	"fmt"
	"math"
)

func bakcPack01(m int, w []int, v []int) int {
	dp := make([]int, m+1)

	for i := 0; i < len(w); i++ {
		for j := m; j >= w[i]; j-- {
			dp[j] = int(math.Max(float64(dp[j]), float64(dp[j-w[i]]+v[i])))
		}
	}

	fmt.Printf("max value: %d\n", dp[m])
	return dp[m]
}

func backPackFull(m int, w []int, v []int) int {
	dp := make([]int, m+1)

	for i := 0; i <= m; i++ {
		for k := 0; k < len(w); k++ {
			if i >= w[k] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[i-w[k]]+v[k])))
			}
		}
	}

	fmt.Printf("max value: %d\n", dp[m])
	return dp[m]
}
