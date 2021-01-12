package sqrt

import "math"

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

//Runtime: 0 ms, faster than 100% of Go online submissions for String to Integer (atoi).
//Memory Usage: 2.3 MB, less than 100% of Go online submissions for String to Integer (atoi).
