package backpack

import (
	"fmt"
)

func bakcPack01(m int, w []int, v []int) int {
	fmt.Printf("01 backpack m: %d, w: %v, v: %v\n", m, w, v)
	dp := make([]int, m+1)

	for i := 0; i < len(w); i++ {
		for j := m; j >= w[i]; j-- {
			dp[j] = max(dp[j], dp[j-w[i]]+v[i])
		}
	}

	fmt.Printf("max value: %d\n", dp[m])
	return dp[m]
}

func backPackFull(m int, w []int, v []int) int {
	fmt.Printf("full backpack m: %d, w: %v, v: %v\n", m, w, v)
	dp := make([]int, m+1)

	for i := 0; i <= m; i++ {
		for k := 0; k < len(w); k++ {
			if i >= w[k] {
				dp[i] = max(dp[i], dp[i-w[k]]+v[k])
			}
		}
	}

	fmt.Printf("max value: %d\n", dp[m])
	return dp[m]
}
