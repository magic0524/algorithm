package bracket

import "fmt"

func bracketMax(s string) int {
	stack := make([]int, 0)
	maxNum := 0
	curNum := 0

	defer func() {
		fmt.Printf("========================\n")
		fmt.Printf("s: %s\tmaxNum: %d\n", s, maxNum)
	}()

	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				curNum = 0
				continue
			} else {
				stack = stack[:len(stack)-1]
				curNum += 2
				maxNum = max(maxNum, curNum)
			}
		}
	}

	return maxNum
}
