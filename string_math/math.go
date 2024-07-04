package stringmath

import "fmt"

func math(a string, b string) int {
	dp := make([][]int, len(a)+1)

	defer func() {
		fmt.Printf("a %s\tb %s\tmax len %d\n", a, b, dp[len(a)][len(b)])
	}()

	for i := 0; i < len(a)+1; i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 0; i < len(a)+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < len(b)+1; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < len(b)+1; j++ {
			if a[i-1] != b[j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			}
		}
	}

	return dp[len(a)][len(b)]
}
