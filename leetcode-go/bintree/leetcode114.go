package bintree

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/
func flatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right
	root.Right = left
	root.Left = nil
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
	return root
}
