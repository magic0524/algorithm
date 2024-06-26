package bracket

import "fmt"

func bracketMax(s string) int {
	fmt.Printf("========================\n")
	stack := make([]int, 0)
	maxAns := 0

	defer func() {
		fmt.Printf("s: %s\tmaxAns: %d\n", s, maxAns)
	}()

	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
		fmt.Printf("i: %d\tstack: %v\n", i, stack)
	}

	return maxAns
}
