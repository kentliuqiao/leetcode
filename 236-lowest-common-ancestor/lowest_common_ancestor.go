package lowestcommonancestor

type TreeNode struct {
	val         int
	left, right *TreeNode
}

/**

前提条件：
1. 树中无重复节点；
2. p，q均存在于树中；
3. p != q。

解题思路：
符合条件的最近公共祖先x满足以下条件之一即可：
1. x的左右子树均包含p或q，则x就是其最近公共祖先；
2. x恰好是p，q其中一个节点，且x的左右子树中存在另外一个节点。

*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.val == p.val || root.val == q.val {
		// 当前树包含p或q节点
		// return root
		return &TreeNode{val: root.val} // 可以直接返回root，新建一个节点仅仅为了方便测试
	}
	// 注意：必须使用后序遍历
	// traverse tree in post order so that we start from the bottom
	left := lowestCommonAncestor(root.left, p, q)
	right := lowestCommonAncestor(root.right, p, q)
	if left != nil && right != nil {
		// 左右子树都有，则返回root
		// return root
		return &TreeNode{val: root.val}
	}
	if left != nil {
		// return left
		return &TreeNode{val: left.val}
	}
	// return right
	if right != nil {
		return &TreeNode{val: right.val}
	}
	return nil
}

func lowestCommonAncestorMap(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	visited := make(map[int]bool)

	// 之所以能够使用这种方法建立每个节点和其父节点的映射关系，在于树中无重复节点！
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.left != nil {
			parent[node.left.val] = node
			dfs(node.left)
		}
		if node.right != nil {
			parent[node.right.val] = node
			dfs(node.right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.val] = true
		p = parent[p.val]
	}
	for q != nil {
		if visited[q.val] {
			return &TreeNode{val: q.val}
		}
		q = parent[q.val]
	}
	return nil
}
