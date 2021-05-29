package allpathsinbinarytree

import "strconv"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	var constructPath func(*TreeNode, string)
	constructPath = func(node *TreeNode, s string) {
		if node == nil {
			return
		}
		path := s + strconv.Itoa(node.val)
		if node.left == nil && node.right == nil {
			paths = append(paths, path)
		} else {
			path += "->"
			constructPath(node.left, path)
			constructPath(node.right, path)
		}
	}
	constructPath(root, "")

	return paths
}
