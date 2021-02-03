package divideTwoIntegers

import "math"

// https://leetcode.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	r:=dividend/divisor
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	return r
}

//Runtime: 4 ms, faster than 74.18% of Go online submissions for Divide Two Integers.
//Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Divide Two Integers.