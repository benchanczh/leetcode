package main

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	for _, node := range root.Children {
		result = append(result, preorder(node)...)
	}
	result = append([]int{root.Val}, result...)

	return result
}
