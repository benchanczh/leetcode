package main

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else if node1.Val != node2.Val {
		return false
	}

	return isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}
