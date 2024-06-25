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
