package winmid

import (
	"fmt"
)

func winMid(arr []int, k int) []int {
	fmt.Printf("====================\n")
	ans := []int{}

	defer func() {
		fmt.Printf("arr %v, k %v, ans %v\n", arr, k, ans)
	}()

	if len(arr) < 2 {
		ans = append(ans, arr[0])
		return ans
	}

	mid := 0
	maxHeap := []item{}
	minHeap := []item{}

	for i := 0; i < k; i++ {
		maxHeap = maxHeapAppend(maxHeap, item{arr[i], i})
	}
	for i := 0; i < k/2; i++ {
		var item item
		item, maxHeap = maxHeapPop(maxHeap)
		minHeap = minHeapAppend(minHeap, item)
	}

	if k%2 == 0 {
		mid = (maxHeap[0].value + minHeap[0].value) / 2
	} else {
		mid = maxHeap[0].value
	}
	ans = append(ans, mid)

	for i := k; i < len(arr); i++ {
		if arr[i] < mid {
			maxHeap = maxHeapAppend(maxHeap, item{arr[i], i})
			var item item
			item, maxHeap = maxHeapPop(maxHeap)
			minHeap = minHeapAppend(minHeap, item)

		} else {
			minHeap = minHeapAppend(minHeap, item{arr[i], i})
			var item item
			item, minHeap = minHeapPop(minHeap)
			maxHeap = maxHeapAppend(maxHeap, item)

		}

		if maxHeap[0].index == i-k {
			_, maxHeap = maxHeapPop(maxHeap)
		} else if minHeap[0].index == i-k {
			_, minHeap = minHeapPop(minHeap)
		}

		if k%2 == 0 {
			mid = (maxHeap[0].value + minHeap[0].value) / 2
		} else {
			mid = maxHeap[0].value
		}

		ans = append(ans, mid)
	}

	return ans
}
