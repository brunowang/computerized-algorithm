package bintree

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	maxDepthWithCalc(root, &maxDiameter)
	return maxDiameter
}

func maxDepthWithCalc(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}
	lmax := maxDepthWithCalc(root.Left, maxDiameter)
	rmax := maxDepthWithCalc(root.Right, maxDiameter)
	*maxDiameter = max(*maxDiameter, lmax+rmax)
	return max(lmax, rmax) + 1
}
