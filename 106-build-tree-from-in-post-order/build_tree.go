package buildtreefrominpostorder

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func buildTree(inOrder, postOrder []int) *TreeNode {
	idxMap := make(map[int]int)
	for k, v := range inOrder {
		idxMap[v] = k
	}
	var build func(int, int) *TreeNode
	build = func(inOrderLeft, inOrderRight int) *TreeNode {
		if inOrderLeft > inOrderRight {
			return nil
		}
		val := postOrder[len(postOrder)-1]
		postOrder = postOrder[:len(postOrder)-1]
		root := &TreeNode{val: val}

		inOrderIdx := idxMap[val]
		root.right = build(inOrderIdx+1, inOrderRight)
		root.left = build(inOrderLeft, inOrderIdx-1)
		return root
	}
	return build(0, len(inOrder)-1)
}
