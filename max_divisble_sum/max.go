package main

import "fmt"

func maxSum(arr []int) int {
	fmt.Printf("arr = %v\n", arr)
	dp := make([][3]int, len(arr))

	n := arr[0] % 3
	dp[0][n] = arr[0]

	for i := 1; i < len(arr); i++ {
		for j := 0; j < 3; j++ {
			sum := dp[i-1][j] + arr[i]
			m := sum % 3
			dp[i][m] = max(dp[i][m], sum)
		}

		for j := 0; j < 3; j++ {
			dp[i][j] = max(dp[i][j], dp[i-1][j])
		}
	}

	fmt.Printf("final dp = %v\n", dp)

	return dp[len(arr)-1][0]
}
