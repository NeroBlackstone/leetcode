package sameTree

import "github.com/NeroBlackstone/leetcode/collection"

// https://leetcode.com/problems/same-tree/
func isSameTree(p *collection.TreeNode, q *collection.TreeNode) bool {
	// p and q are both null
	if p == nil && q == nil {
		return true
	}
	// one of p and q is null
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
//Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Same Tree.
// 总结：记住这种二叉树的递归比较方法。
