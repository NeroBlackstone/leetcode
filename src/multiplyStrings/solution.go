package multiplyStrings

import "math/big"

// https://leetcode.com/problems/multiply-strings/
func multiply(num1 string, num2 string) string {
	var n1,n2,res big.Int
	n1.SetString(num1,10)
	n2.SetString(num2,10)
	res.Mul(&n1,&n2)
	return res.Text(10)
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Multiply Strings.
//Memory Usage: 2.6 MB, less than 56.07% of Go online submissions for Multiply Strings.