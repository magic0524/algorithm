package stock

import "fmt"

func stockMax(arr []int) int {
	fmt.Printf("=========== stockMax ===========\n")

	maxNum := 0
	defer func() {
		fmt.Printf("arr = %v\tmaxNum = %v\n", arr, maxNum)
	}()

	stack := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if len(stack) > 0 && arr[i] < stack[len(stack)-1] {
			n := stack[len(stack)-1] - stack[0]
			maxNum += n
			stack = stack[:0]
		}

		stack = append(stack, arr[i])
	}

	if len(stack) > 0 {
		n := stack[len(stack)-1] - stack[0]
		maxNum += n
	}

	return maxNum
}

func stockMaxDP(arr []int) int {
	maxNum := 0
	defer func() {
		fmt.Printf("=========== stockMaxDP ===========\n")
		fmt.Printf("arr = %v\tmaxNum = %v\n", arr, maxNum)
	}()

	dp := make([][2]int, len(arr))
	dp[0][0] = 0
	dp[0][1] = -arr[0]

	for i := 1; i < len(arr); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+arr[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-arr[i])
	}

	maxNum = dp[len(arr)-1][0]

	return maxNum
}

func stockMaxOneBuyDP(arr []int) int {
	maxNum := 0
	defer func() {
		fmt.Printf("=========== stockMaxOneBuyDP ===========\n")
		fmt.Printf("arr = %v\tmaxNum = %v\n", arr, maxNum)
	}()

	dp := make([][2]int, len(arr))
	dp[0][0] = 0
	dp[0][1] = -arr[0]

	for i := 1; i < len(arr); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+arr[i])
		dp[i][1] = max(dp[i-1][1], -arr[i])
	}

	maxNum = dp[len(arr)-1][0]

	return maxNum
}
