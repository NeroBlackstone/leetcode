package main

func main() {
	a := TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}
	inorderTraversal(&a)
}

// leetcode 94 中序遍历二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 第二次默写完成
func inorderTraversal(root *TreeNode) []int {
	var res []int
	// 开始操作
	//定义个函数先
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		// 空值返回
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	// 完成
	return res
}
