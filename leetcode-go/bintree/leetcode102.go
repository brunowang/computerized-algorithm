package bintree

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		level := make([]int, 0)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
