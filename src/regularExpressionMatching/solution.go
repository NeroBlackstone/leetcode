package regularExpressionMatching

//'.' Matches any single character.
//'*' Matches zero or more of the preceding element.
//s could be empty and contains only lowercase letters a-z.
//p could be empty and contains only lowercase letters a-z, and characters like . or *.

// 思路：使用pattern字符串分块匹配，两个两个字符去读取匹配，检测到后一个是*就保留这俩，不是*就只去匹配前一个

func isMatch(s string, p string) bool {
	sIndex:=0
	pIndex:=0
	sLength:=len(s)
	pLength:=len(p)

	//precedingRune:=0

	// 对每个pattern的字符，比对原string
	for i,c:=range p{
		// 在golang中，字符串本质是8-bit bytes的只读切片，但是for range的时候并不是去拿单个bytes值，而是去拿整个rune字符。
		// 所以对比字符的时候要强制对原字符串转rune切片，再对比单个rune
		if i!=
	}
}