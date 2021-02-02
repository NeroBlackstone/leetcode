package numberOf1Bits

import "math/bits"

// https://leetcode.com/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
    return bits.OnesCount32(num)
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of 1 Bits.
//Memory Usage: 2 MB, less than 8.46% of Go online submissions for Number of 1 Bits.