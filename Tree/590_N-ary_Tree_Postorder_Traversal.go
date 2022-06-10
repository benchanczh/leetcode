package main

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	for _, node := range root.Children {
		result = append(result, postorder(node)...)
	}
	result = append(result, root.Val)

	return result
}
