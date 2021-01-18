package integerToRoman

func intToRoman(num int) string {
	type table struct {
		value  int
		symbol string
	}
	t := []table{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	s := ""
	for _, e := range t {
		v := e.value
		for v <= num {
			num -= v
			s += e.symbol
		}
		if num == 0 {
			break
		}
	}
	return s
}

// Runtime: 4 ms, faster than 93.75% of Go online submissions for Integer to Roman.
// Memory Usage: 3.5 MB, less than 27.73% of Go online submissions for Integer to Roman.
// 总结：贪心算法
