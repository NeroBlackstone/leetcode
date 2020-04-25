package zigZagConversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	//处理边界情况
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	//斜线至多有numSpace个字母
	numOblique := numRows - 2
	v := numOblique + numRows
	//有n个 ｜/ 字型字符组合
	n := len(s) / v
	//最后蛇尾rest个字符
	rest := len(s) % v
	//竖线组成的字符串数组切片
	stringRows := make([]string, n+1)
	//斜线组成的字符串数组切片
	stringOblique := make([]string, n+1)
	//每次读numRows个字符到字节切片中
	byteRows := make([]byte, numRows)
	byteOblique := make([]byte, numOblique)
	reader := strings.NewReader(s)
	//循环读取字符串到两个字符串数组切片内
	for i := 0; ; i++ {
		//实际读了nRows/nObliques个字符
		//注意这里有一个坑，第一次Read读到末尾，只要读到了字符串，就不会抛出EOF，只有完全没读到字符串时，才会抛出EOF
		nRows, _ := reader.Read(byteRows)
		nObliques, _ := reader.Read(byteOblique)
		//从[]byte缓存切片内拿出数据到字符串数组里
		stringRows[i] = string(byteRows[:nRows])
		stringOblique[i] = string(byteOblique[:nObliques])
		//字符串读完了
		if reader.Len() == 0 {
			break
		}
	}
	var builder strings.Builder
	hasOblique := false
	restOblique := 0
	if len(stringOblique[n]) != 0 {
		hasOblique = true
	}
	if hasOblique {
		restOblique = rest - numRows
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < n+1; j++ {
			//防止越界
			if i < len(stringRows[j]) {
				builder.WriteByte(stringRows[j][i])
			}
			//非第一行和最后一行
			if i > 0 && i < numRows-1 {
				//输出中间斜线，不在最后的情况
				if j != n {
					//逆序输出
					builder.WriteByte(stringOblique[j][numOblique-i])
					//在最后且有尾部的情况
				} else if hasOblique {
					//确定从哪一行开始计算
					if i > numOblique-restOblique {
						//逆序输出
						builder.WriteByte(stringOblique[j][numOblique-i])
					}
				}
			}
		}
	}
	return builder.String()
}

//Runtime: 4 ms, faster than 92.78% of Go online submissions for ZigZag Conversion.
//Memory Usage: 4.4 MB, less than 71.43% of Go online submissions for ZigZag Conversion.
