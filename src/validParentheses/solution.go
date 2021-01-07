package validParentheses

import (
	"fmt"
	"github.com/NeroBlackstone/leetcode/collection"
)

func isValid(s string) bool {
	sLength := len(s)
	// 1 <= s.length <= 10^4
	if sLength > 10000 || sLength < 1 {
		return false
	}

	stack := collection.InitStack()

	for _, r := range s {
		// s consists of parentheses only '()[]{}'.
		if r != '(' && r != ')' && r != '[' && r != ']' && r != '{' && r != '}' {
			return false
		} else {
			if r == '(' || r == '[' || r == '{' {
				stack.Push(r)
			} else {
				top, err := stack.Pop()
				if err != nil {
					fmt.Println(err)
					return false
				}
				l := top.(int32)
				if (l == '(' && r != ')') || (l == '{' && r != '}') || (l == '[' && r != ']') {
					return false
				}
			}
		}
	}
	if !stack.IsEmpty() {
		return false
	}
	return true
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.2 MB, less than 21.55% of Go online submissions for Valid Parentheses.
// 总结：很简单的Stack实现
