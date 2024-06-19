package warmweather

import "fmt"

// 使用单调栈解决
func warmWeather(a []int) []int {
	fmt.Printf("day weather = %v\n", a)
	stack := make([]int, 0)
	ans := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		for len(stack) > 0 && a[stack[len(stack)-1]] < a[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}

		stack = append(stack, i)
	}

	fmt.Printf("ans = %v\n", ans)

	return ans
}
