package palindromeNumber

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

//Runtime: 20 ms, faster than 41.17% of Go online submissions for Palindrome Number.
//Memory Usage: 5.4 MB, less than 25.00% of Go online submissions for Palindrome Number.

//总结：太容易了。。。
