package main

import "fmt"

func lis(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}

	dp := make([]int, len(a))
	dp[0] = 1

	maxLen := 0

	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	fmt.Printf("maxLen = %d\n", maxLen)
	return maxLen
}
