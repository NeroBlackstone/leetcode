package letterCombinationsOfAPhoneNumber

import "fmt"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	m := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	output := make([]string, 0)

	for _, d := range digits {
		tempSlice := make([]string, 0)
		for _, r := range m[d] {
			if len(output) == 0 {
				tempSlice = append(tempSlice, string(r))
			} else {
				for _, s := range output {
					tempSlice = append(tempSlice, s+string(r))
				}
			}
		}
		output = tempSlice
		fmt.Println(output)
	}
	return output
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
// Memory Usage: 2.2 MB, less than 10.33% of Go online submissions for Letter Combinations of a Phone Number.
// summary: 其实也可以用回溯法（backtracking），但是我觉得golang的语言设计不适合在这里使用回溯。
