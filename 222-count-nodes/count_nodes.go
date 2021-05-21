package countnodes

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func countNodes(root *TreeNode) (res int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.right == nil || prev == root.right {
			res++
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.right
		}
	}
	return
}

func countNodesPreOrderRecursive(root *TreeNode) (res int) {
	var preOrder func(*TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res++
		preOrder(node.left)
		preOrder(node.right)
	}
	preOrder(root)

	return
}

// 规定根节点位于第 0 层，完全二叉树的最大层数为 h。
// 对于最大层数为 h 的完全二叉树，节点个数一定在 [2^h,2^(h+1)-1] 的范围内，
// 可以在该范围内通过二分查找的方式得到完全二叉树的节点个数。
/*

具体做法是，根据节点个数范围的上下界得到当前需要判断的节点个数 k，
如果第 k 个节点存在，则节点个数一定大于或等于 k，
如果第 k 个节点不存在，则节点个数一定小于 k，
由此可以将查找的范围缩小一半，直到得到节点个数。

如何判断第 k 个节点是否存在呢？
如果第 k 个节点位于第 h 层，则 k 的二进制表示包含 h+1 位，其中最高位是 1，
其余各位从高到低表示从根节点到第 k 个节点的路径，0 表示移动到左子节点，1 表示移动到右子节点。
通过位运算得到第 k 个节点对应的路径，判断该路径对应的节点是否存在，即可判断第 k 个节点是否存在。

*/
func countNodesOptimization(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}
	height := 0
	for node := root; node != nil; {
		height++
		node = node.left
	}
	min, max := 1<<height, (1<<(height+1))-1
	for min < max {
		mid := min + (max-min+1)/2
		if exists(root, height, mid) {
			min = mid
		} else {
			max = mid - 1
		}
	}
	return min
}

func exists(root *TreeNode, height, k int) bool {
	bits := 1 << (height - 1)
	node := root
	for node != nil && bits > 0 {
		if bits&k == 0 {
			node = node.left
		} else {
			node = node.right
		}
		bits >>= 1
	}
	return node != nil
}
