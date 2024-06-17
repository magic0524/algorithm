package ringarr

import "fmt"

func max(a []int) []int {
	fmt.Printf("a = %v\n", a)

	stack := make([]int, 0)
	ans := make([]int, len(a))
	for i := range ans {
		ans[i] = -1
	}

	for k := 0; k < len(a)*2; k++ {
		i := k % len(a)
		for len(stack) > 0 && a[i] > a[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = a[i]
		}

		stack = append(stack, i)
	}

	fmt.Printf("ans: %v\n", ans)

	return ans
}
