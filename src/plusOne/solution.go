package plusOne

// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	lastIs9 := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			lastIs9 = false
			break
		} else {
			digits[i] = 0
			lastIs9 = true
		}
	}
	if lastIs9 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
// Memory Usage: 2 MB, less than 32.90% of Go online submissions for Plus One.
