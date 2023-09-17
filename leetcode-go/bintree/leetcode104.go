package bintree

/*
给定一个二叉树 root ，返回其最大深度。
二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lmax := maxDepth(root.Left)
	rmax := maxDepth(root.Right)
	return max(lmax, rmax) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
