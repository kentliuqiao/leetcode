package allpathsintree

import "strconv"

var paths []string

type TreeNode struct {
	val         int
	left, right *TreeNode
}

// BFS
func allPaths(root *TreeNode) []string {
	paths = []string{}
	constructPath(root, "")
	return paths
}

func constructPath(node *TreeNode, path string) {
	if node != nil {
		pathSB := path + strconv.Itoa(node.val)
		if node.left == nil && node.right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			constructPath(node.left, pathSB)
			constructPath(node.right, pathSB)
		}
	}
}

// DFS
func allPathsRecursive(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	nodeQ, pathQ := []*TreeNode{root}, []string{strconv.Itoa(root.val)}

	for len(nodeQ) > 0 {
		node := nodeQ[0]
		nodeQ = nodeQ[1:]
		path := pathQ[0]
		pathQ = pathQ[1:]
		if node.left == nil && node.right == nil {
			paths = append(paths, path)
			continue
		}
		if node.left != nil {
			nodeQ = append(nodeQ, node.left)
			pathQ = append(pathQ, path+"->"+strconv.Itoa(node.left.val))
		}
		if node.right != nil {
			nodeQ = append(nodeQ, node.right)
			pathQ = append(pathQ, path+"->"+strconv.Itoa(node.right.val))
		}
	}
	return paths
}
