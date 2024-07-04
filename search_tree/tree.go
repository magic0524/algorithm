package searchtree

func searchTree(arr []int, a, b, c int) int {
	if len(arr) == 0 {
		return -1
	}

	if a < arr[0] && b < arr[0] && c < arr[0] {
		return searchTree(arr[1:], a, b, c)
	} else if a > arr[0] && b > arr[0] && c > arr[0] {
		return searchTree(arr[2:], a, b, c)
	} else {
		return arr[0]
	}
}
