package buildtreefrompreinorder

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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
		root := &TreeNode{Val: rootVal}
		preOrder = preOrder[1:]
		inOrderIdx := idxMap[rootVal]
		root.Left = build(inOrderLeft, inOrderIdx-1)
		root.Right = build(inOrderIdx+1, inOrderRight)
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
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func build(preorder []int, prelo, prehi int, inorder []int, inlo, inhi int) *TreeNode {
	if len(preorder) == 0 || prelo > prehi {
		return nil
	}
	root := &TreeNode{Val: preorder[prelo]}

	index := -1
	for i := inlo; i <= inhi; i++ {
		if inorder[i] == preorder[prelo] {
			index = i
			break
		}
	}
	if index == -1 {
		return root
	}
	left := build(preorder, prelo+1, prelo+index-inlo, inorder, inlo, index-1)
	right := build(preorder, prelo+index-inlo+1, prehi, inorder, index+1, inhi)

	root.Left = left
	root.Right = right

	return root
}
