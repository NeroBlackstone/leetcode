package romanToInteger

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	// validate string
	// length of String
	sLength := len(s)
	// 1 <= s.length <= 15
	if sLength < 1 || sLength > 15 {
		return 0
	}
	// s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
	for _, v := range s {
		if v != 'I' && v != 'V' && v != 'X' && v != 'L' && v != 'C' && v != 'D' && v != 'M' {
			return 0
		}
	}
	// validate string end

	// init map for String to integer
	m := map[string]int{
		"I":  1,
		"II": 2,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	// return integer
	var integer int

	for i := 0; i < sLength; {
		if i < sLength-1 {
			if v, ok := m[s[i:i+2]]; ok {
				integer += v
				i += 2
			} else {
				integer += m[s[i:i+1]]
				i++
			}
		} else {
			integer += m[s[i:i+1]]
			i++
		}
	}

	// It is guaranteed that s is a valid roman numeral in the range [1, 3999].
	if integer < 1 || integer > 3999 {
		return 0
	}
	return integer
}

//Runtime: 12 ms, faster than 30.70% of Go online submissions for Roman to Integer.
//Memory Usage: 6.1 MB, less than 10.98% of Go online submissions for Roman to Integer.
//总结：该题目的关键是识别出让数字减少的三种情况。更好的方法是从右到左读取字符串。如果该字符串是IXC这种可以减少数字大小的字符，
//且已累计的数字大于该字符串让其他字符串减少的本身实际表义，则说明它是减少后面字符大小的作用。否则直接加上字符本身表义即可。
