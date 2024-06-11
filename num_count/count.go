package main

import "fmt"

func main() {
	a := []int{5, 1, 5, 8, 3, 3, 6, 9, 2, 7, 7, 7, 7, 7}
	fmt.Printf("Before: %v\n", a)

	num(a)
}

func num(a []int) []int {
	for i := 0; i < len(a); {
		if a[i] <= 0 {
			i++
		} else {
			idx := a[i]
			if a[idx] == 0 {
				a[idx] = -1
				a[i] = 0
				i++
			} else if a[idx] < 0 {
				a[idx]--
				a[i] = 0
				i++
			} else {
				if idx == i {
					a[i] = -1
				} else {
					b := a[idx]
					a[idx] = -1
					a[i] = b
				}
			}
		}
	}

	for i, v := range a {
		if v < 0 {
			a[i] = -v
		}
	}

	fmt.Printf("After: %v\n", a)

	return a
}
