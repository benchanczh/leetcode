package main

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
