package bintree

/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/
func preorderTraversal(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)
	return ret
}
