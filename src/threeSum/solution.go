package threeSum

import (
	"sort"
)

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i, e1 := range nums {
		if i > 0 && e1 == nums[i-1] {
			continue
		}
		restNums := nums[i+1:]
		k := len(restNums) - 1
		sum := -1 * e1
		for j, e2 := range restNums {
			if j > 0 && e2 == restNums[j-1] {
				continue
			}
			for j < k && e2+restNums[k] > sum {
				k--
			}
			if j == k {
				break
			}
			e3 := restNums[k]
			if e2+e3 == sum {
				ans = append(ans, []int{e1, e2, e3})
			}
		}
	}
	return ans
}

// Runtime: 28 ms, faster than 96.44% of Go online submissions for 3Sum.
// Memory Usage: 7 MB, less than 42.28% of Go online submissions for 3Sum.
// Summary : two point
