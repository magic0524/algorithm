package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 1, 5, 6, 4, 10, 3, 5, 9, 3, 10}
	fmt.Printf("before: %v\n", a)

	fmt.Printf("kth largest: %v\n", findKthLargest(a, 3))
}

func findKthLargest(nums []int, k int) []int {
	// 1. 先构建一个大小为k的小顶堆
	heap := nums[:k]
	buildHeap(heap)

	// 2. 遍历剩余的元素，如果比堆顶元素大，则替换堆顶元素，然后调整堆
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			heapify(heap, 0, k)
		}
	}

	return heap
}

func buildHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i, len(nums))
	}
}

func heapify(nums []int, i, n int) {
	left, right, min := 2*i+1, 2*i+2, i
	if left < n && nums[left] < nums[min] {
		min = left
	}
	if right < n && nums[right] < nums[min] {
		min = right
	}

	if min != i {
		nums[i], nums[min] = nums[min], nums[i]
		heapify(nums, min, n)
	}
}
