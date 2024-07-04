package winmid

type item struct {
	value int
	index int
}

func buildMaxHeap(a []item) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		heapifyMax(a, i, len(a))
	}
}

func heapifyMax(a []item, i int, n int) {
	left, right, max := 2*i+1, 2*i+2, i

	if left < len(a) && a[left].value > a[max].value {
		max = left
	}

	if right < len(a) && a[right].value > a[max].value {
		max = right
	}

	if max != i {
		a[i], a[max] = a[max], a[i]
		heapifyMax(a, max, n)
	}
}

func maxHeapAppend(a []item, v item) []item {
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

func maxHeapPop(a []item) (item, []item) {
	v := a[0]
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	heapifyMax(a, 0, len(a))
	return v, a
}
