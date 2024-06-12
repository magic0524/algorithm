package lcs

import "math"

func longestCommonSubstring(a, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	// Create a 2D slice to store the length of the longest common substring
	// between the first i characters of a and the first j characters of b
	lcs := make([][]int, len(a)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(b)+1)
	}

	// Initialize the first row and column of the 2D slice to 0
	for i := 0; i <= len(a); i++ {
		lcs[i][0] = 0
	}
	for j := 0; j <= len(b); j++ {
		lcs[0][j] = 0
	}

	max := 0
	// Fill in the 2D slice using dynamic programming
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = 0
			}

			if lcs[i][j] > max {
				max = lcs[i][j]
			}
		}
	}

	return max
}

func longestCommonSubsequnece(a, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	// Create a 2D slice to store the length of the longest common subsequence
	// between the first i characters of a and the first j characters of b
	lcs := make([][]int, len(a)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(b)+1)
	}

	// Initialize the first row and column of the 2D slice to 0
	for i := 0; i <= len(a); i++ {
		lcs[i][0] = 0
	}
	for j := 0; j <= len(b); j++ {
		lcs[0][j] = 0
	}

	// Fill in the 2D slice using dynamic programming
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = int(math.Max(float64(lcs[i-1][j]), float64(lcs[i][j-1])))
			}
		}
	}

	return lcs[len(a)][len(b)]
}
