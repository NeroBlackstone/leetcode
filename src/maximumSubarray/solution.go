package maximumSubarray

import "github.com/NeroBlackstone/leetcode/collection"

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	maxSum, sum := nums[0], nums[0]
	for _, e := range nums[1:] {
		sum = collection.Max(e, sum+e)
		maxSum = collection.Max(sum, maxSum)
	}
	return maxSum
}

// Runtime: 4 ms, faster than 96.28% of Go online submissions for Maximum Subarray.
// Memory Usage: 3.3 MB, less than 100.00% of Go online submissions for Maximum Subarray.
// 总结：熟悉Kadane’s Algorithm
