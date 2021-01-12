package addBinary

import "math/big"

//https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	aInt, bInt, sum := new(big.Int), new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)
	sum.Add(aInt, bInt)
	return sum.Text(2)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
//Memory Usage: 2.4 MB, less than 63.44% of Go online submissions for Add Binary.
// 总结：要熟悉"math/big"包的使用
