package constructbinarytreefromprepostorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func build(preorder []int, prelo, prehi int, postorder []int, postlo, posthi int) *TreeNode {
	if prelo > prehi {
		return nil
	}
	if prelo == prehi {
		return &TreeNode{Val: preorder[prehi]}
	}

	rootVal := preorder[prelo]
	index := -1
	for i := postlo; i < posthi; i++ {
		if postorder[i] == preorder[prelo+1] {
			index = i
			break
		}
	}
	leftSize := index - postlo
	root := &TreeNode{Val: rootVal}
	left := build(preorder, prelo+1, prelo+1+leftSize, postorder, postlo, index)
	right := build(preorder, prelo+2+leftSize, prehi, postorder, index+1, posthi-1)
	root.Left = left
	root.Right = right
	return root
}
