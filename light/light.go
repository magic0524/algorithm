package main

import (
	"math"
)

// 第n个数，首先一定会被1和n这两次开关操作，抵消了，是关的
// n如果能被整除，说明n=x*y，x!=y，也就是会被x和y两次操作，同样抵消了，还是关的
// n如果能不能被整除，说明不会被影响
// n如果被开方，n=z^2，只会被z影响一次，是开的
func light(n int) int {
	return int(math.Sqrt(float64(n)))
}
