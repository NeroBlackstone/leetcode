package removeElement

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	length := 0
	for i, e := range nums {
		if e != val {
			nums[length] = nums[i]
			length++
		}
	}
	nums = nums[:length]
	return length
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
//Memory Usage: 2.1 MB, less than 9.47% of Go online submissions for Remove Element.
// 总结：又是一道双指针题
