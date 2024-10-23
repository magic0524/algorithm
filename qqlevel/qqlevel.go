package main

import "fmt"

func main() {
	var groupNum int
	fmt.Scanln(&groupNum)
	for i := 0; i < groupNum; i++ {
		m, s, c, level := 0, 0, 0, 0
		fmt.Scanln(&m, &s, &c)
		fmt.Scanln(&level)
		fmt.Println(minIcon(m, s, c, level))
	}
}

func minIcon(m, s, c, level int) int {
	if m <= 0 || s <= 0 || c <= 0 || level <= 0 {
		return 0
	}

	starM := m
	starS := m * s
	starC := m * s * c

	cCount := level / starC
	less := level % starC

	sCount := less / starS
	less = less % starS

	mCount := less / starM
	less = less % starM

	return cCount + sCount + mCount + less
}
