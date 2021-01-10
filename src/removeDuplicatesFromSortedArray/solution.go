package removeDuplicatesFromSortedArray

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	amount := 0
	for _, e := range nums {
		if e != nums[amount] {
			amount++
			nums[amount] = e
		}
	}
	length := amount + 1
	nums = nums[:length]
	return length
}

// Runtime: 12 ms, faster than 26.35% of Go online submissions for Remove Duplicates from Sorted Array.
//Memory Usage: 4.6 MB, less than 9.56% of Go online submissions for Remove Duplicates from Sorted Array.
// 总结： 非常经典的双指针思想
