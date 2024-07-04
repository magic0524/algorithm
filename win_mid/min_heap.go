package winmid

func buildMinHeap(a []item) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		heapifyMin(a, i, len(a))
	}
}

func heapifyMin(a []item, i int, n int) {
	left, right, min := 2*i+1, 2*i+2, i

	if left < len(a) && a[left].value < a[min].value {
		min = left
	}

	if right < len(a) && a[right].value < a[min].value {
		min = right
	}

	if min != i {
		a[i], a[min] = a[min], a[i]
		heapifyMin(a, min, n)
	}
}

func minHeapAppend(a []item, v item) []item {
	a = append(a, v)
	i := len(a) - 1
	for i > 0 {
		j := (i - 1) / 2
		if a[j].value <= v.value {
			break
		}
		a[i] = a[j]
		i = j
	}
	a[i] = v
	return a
}

func minHeapPop(a []item) (item, []item) {
	v := a[0]
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	heapifyMin(a, 0, len(a))
	return v, a
}
