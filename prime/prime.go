package prime

import (
	"fmt"
	"math"
)

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
