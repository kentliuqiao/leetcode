package deletenodeinbst

type TreeNode struct {
	val         int
	left, right *TreeNode
}

// func deleteNodeInBST(root *TreeNode, target int) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	node := root
// 	for node != nil {
// 		if node.val == target {
// 			var newRoot *TreeNode
// 			left := node.left
// 			for left != nil {
// 				newRoot = node.left
// 				left = node.left
// 			}
// 			if newRoot == nil {

// 			}
// 		}
// 	}
// }
