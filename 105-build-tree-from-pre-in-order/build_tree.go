package buildtreefrompreinorder

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func buildTreeV2(preOrder, inOrder []int) *TreeNode {
	idxMap := make(map[int]int)
	for k, v := range inOrder {
		idxMap[v] = k
	}
	var build func(int, int) *TreeNode
	build = func(inOrderLeft, inOrderRight int) *TreeNode {
		if inOrderLeft > inOrderRight {
			return nil
		}
		rootVal := preOrder[0]
		root := &TreeNode{val: rootVal}
		preOrder = preOrder[1:]
		inOrderIdx := idxMap[rootVal]
		root.left = build(inOrderLeft, inOrderIdx-1)
		root.right = build(inOrderIdx+1, inOrderRight)
		return root
	}
	return build(0, len(inOrder)-1)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
