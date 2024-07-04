package lps

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))

	maxLen := 1
	start := 0
	defer func() {
		fmt.Printf("s %s\tmaxLen = %d\tsub string %s\n", s, maxLen, s[start:start+maxLen])
	}()
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s)-l+1; i++ {
			j := i + l - 1
			if l == 2 && s[i] == s[j] {
				dp[i][j] = true
				start = i
				maxLen = l
			} else if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				start = i
				maxLen = l
			}
		}
	}
	return s[start : start+maxLen]
}
