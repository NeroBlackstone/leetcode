package threeSumClosest

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	diff, sum := math.MaxInt16, 0
	sort.Ints(nums)
out:
	for i, e := range nums {
		l, h := i+1, len(nums)-1
		for l < h {
			nowSum := e + nums[l] + nums[h]
			nowDiff := target - nowSum
			if math.Abs(float64(nowDiff)) < math.Abs(float64(diff)) {
				diff = nowDiff
				sum = nowSum
				if diff == 0 {
					break out
				}
			}
			if nowSum < target {
				l++
			} else {
				h--
			}
		}
	}
	return sum
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for 3Sum Closest.
// Memory Usage: 2.7 MB, less than 20.73% of Go online submissions for 3Sum Closest.
// summary: two pointer
