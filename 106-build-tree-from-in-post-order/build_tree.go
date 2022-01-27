package buildtreefrominpostorder

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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
		root := &TreeNode{Val: val}

		inOrderIdx := idxMap[val]
		root.Right = build(inOrderIdx+1, inOrderRight)
		root.Left = build(inOrderLeft, inOrderIdx-1)
		return root
	}
	return build(0, len(inOrder)-1)
}

func buildTreeV2(inOrder, postOrder []int) *TreeNode {
	if len(postOrder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postOrder[len(postOrder)-1]}
	i := 0
	for ; i < len(inOrder); i++ {
		if inOrder[i] == root.Val {
			break
		}
	}
	root.Left = buildTreeV2(inOrder[:i], postOrder[:len(inOrder[:i])])
	root.Right = buildTreeV2(inOrder[i+1:], postOrder[len(inOrder[:i]):len(postOrder)-1])
	return root
}

func build(inorder []int, inlo, inhi int, postorder []int, postlo, posthi int) *TreeNode {
	if postlo > posthi {
		return nil
	}

	rootVal := postorder[posthi]
	index := -1
	for i := inlo; i <= inhi; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}

	leftSize := index - inlo
	root := &TreeNode{Val: rootVal}
	left := build(inorder, inlo, index-1, postorder, postlo, postlo+leftSize-1)
	right := build(inorder, index+1, inhi, postorder, postlo+leftSize, posthi-1)
	root.Left = left
	root.Right = right
	return root
}
