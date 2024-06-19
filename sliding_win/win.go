package main

import "fmt"

// 使用最大堆的方法，时间复杂度为 O(nlogk)
func main() {
	a := []int{5, 3, 2, 1, 4, 3, 2, 1, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("a = %v\n", a)

	it := make([]item, len(a))
	for i, v := range a {
		it[i] = item{v, i}
	}
	fmt.Printf("findKthLargest(a, 5) = %v\n", findKthLargest(it, 5))
}

func findKthLargest(a []item, k int) []int {
	ans := make([]int, len(a)-k+1)
	b := a[:k]
	buildHeap(b)

	ans[0] = b[0].value

	for i := k; i < len(a); i++ {
		b = heapAppend(b, a[i])
		for b[0].index < i-k+1 {
			heapPop(b)
		}

		ans[i-k+1] = b[0].value
	}

	return ans
}

type item struct {
	value int
	index int
}

func buildHeap(a []item) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		heapify(a, i, len(a))
	}
}

func heapify(a []item, i int, n int) {
	left, right, max := 2*i+1, 2*i+2, i

	if left < len(a) && a[left].value > a[max].value {
		max = left
	}

	if right < len(a) && a[right].value > a[max].value {
		max = right
	}

	if max != i {
		a[i], a[max] = a[max], a[i]
		heapify(a, max, n)
	}
}

func heapAppend(a []item, v item) []item {
	a = append(a, v)
	i := len(a) - 1
	for i > 0 {
		j := (i - 1) / 2
		if a[j].value >= v.value {
			break
		}
		a[i] = a[j]
		i = j
	}
	a[i] = v
	return a
}

func heapPop(a []item) (item, []item) {
	v := a[0]
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	heapify(a, 0, len(a))
	return v, a
}
