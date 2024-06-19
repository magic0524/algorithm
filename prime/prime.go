package prime

import (
	"fmt"
	"math"
)

// 只要n可以被2到n的平方根之间的任意一个数整除，那么n就不是素数
// 优化：只需要判断2到n的平方根之间的素数是否能整除n
func prime(a int) []int {
	prime := make([]int, 0)

	prime = append(prime, 2)

	for i := 3; i <= a; i++ {
		isPrime := true
		for _, j := range prime {
			if j > int(math.Sqrt(float64(i))) {
				break
			}
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = append(prime, i)
		}
	}

	fmt.Printf("prime: %v\n", prime)

	return prime
}
