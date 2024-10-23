package main

import "fmt"

type link struct {
	a int64
	b int64
}

func main() {
	var groupNum int
	fmt.Scanln(&groupNum)

	for i := 0; i < groupNum; i++ {
		var n, m int
		fmt.Scanln(&n, &m)
		links := make([]link, 0, n)
		for j := 0; j < n; j++ {
			link := link{}
			fmt.Scanln(&link.a, &link.b)
			links = append(links, link)
		}

		friend := make([]int64, 0, m)
		for j := 0; j < m; j++ {
			var a int64
			fmt.Scan(&a)
			friend = append(friend, a)
		}

		fmt.Println(sovle(n, m, links, friend))
	}
}

func sovle(n, m int, links []link, friend []int64) int64 {
	if (m <= 0) || (n <= 0) || (len(links) != n) || (len(friend) != m) {
		return 0
	}

	allFriend := make(map[int64]struct{}, m)
	for _, f := range friend {
		allFriend[f] = struct{}{}
	}

	oneFriend := make(map[int64]int64)

	var nowComment int64
	for _, l := range links {
		aF, a := allFriend[l.a]
		bF, b := allFriend[l.b]

		if a && b {
			nowComment++
			continue
		}

		if a && !b {
			oneFriend[l.b]++
			continue
		}

		if !a && b {
			oneFriend[l.a]++
			continue
		}

		if !a && !b && aF == bF {
			oneFriend[l.a]++
			continue
		}
	}

	var newMaxComment int64
	for _, v := range oneFriend {
		if v > newMaxComment {
			newMaxComment = v
		}
	}

	return nowComment + newMaxComment
}
