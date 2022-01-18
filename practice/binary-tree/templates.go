package binarytree

/**
========================================================================
|                          各种数据结构遍历模版                           |
========================================================================

总结：
1. 只要是递归形式的遍历，都会有一个前序和后序位置，分别在递归之前和之后。

2. 前序位置的代码在刚刚进入一个二叉树节点的时候执行；
后序位置的代码在将要离开一个二叉树节点的时候执行；
中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行。

3. 二叉树的所有问题，就是让你在前中后序位置注入巧妙的代码逻辑，去达到自己的目的。

4. 二叉树题目的递归解法可以分两类思路，第一类是遍历一遍二叉树得出答案，
第二类是通过分解问题计算出答案，这两类思路分别对应着 回溯算法核心框架 和 动态规划核心框架。

*/

//------------- binary tree start -----------------------
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归遍历二叉树
func traverseBinaryTree(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序位置
	traverseBinaryTree(root.Left)
	// 中序位置
	traverseBinaryTree(root.Right)
	// 后序位置
}

// 前序遍历二叉树
// 通过遍历的思路完成
func preOrderRecursive(root *TreeNode) []int {
	res := []int{}
	var preOrderRecursiveFunc func(*TreeNode)
	preOrderRecursiveFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preOrderRecursiveFunc(root.Left)
		preOrderRecursiveFunc(root.Right)
	}
	preOrderRecursiveFunc(root)
	return res
}

// 前序遍历二叉树
// 通过分治的思路完成
// 二叉树前序序列：[ 根 -> 左子树前序序列 -> 右子树前序序列 ]
func preOrderDivideConquer(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preOrderDivideConquer(root.Left)...)
	res = append(res, preOrderDivideConquer(root.Right)...)
	return res
}

// 中序遍历二叉树
// 通过遍历的思路完成
func inOrderRecursive(root *TreeNode) []int {
	res := []int{}
	var inOrderRecursiveFunc func(*TreeNode)
	inOrderRecursiveFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderRecursiveFunc(root.Left)
		res = append(res, root.Val)
		inOrderRecursiveFunc(root.Right)
	}
	inOrderRecursiveFunc(root)
	return res
}

// 中序遍历二叉树
// 通过分治的思路完成
// 二叉树中序序列：[ 左子树中序序列 -> 根 -> 右子树中序序列 ]
func inOrderDivideConquer(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inOrderDivideConquer(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inOrderDivideConquer(root.Right)...)
	return res
}

// 后序遍历二叉树
// 通过遍历的思路完成
func postOrderRecursive(root *TreeNode) []int {
	res := []int{}
	var postOrderRecursiveFunc func(*TreeNode)
	postOrderRecursiveFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrderRecursiveFunc(root.Left)
		postOrderRecursiveFunc(root.Right)
		res = append(res, root.Val)
	}
	postOrderRecursiveFunc(root)
	return res
}

// 后序遍历二叉树
// 通过分治的思路完成
// 二叉树的后序序列：[ 左子树的后序序列 -> 右子树的后序序列 -> 根 ]
func postOrderDivideConquer(root *TreeNode) []int {
	res := []int{}
	res = append(res, postOrderDivideConquer(root.Left)...)
	res = append(res, postOrderDivideConquer(root.Right)...)
	res = append(res, root.Val)
	return res
}

//------------- binary tree end -----------------------

//------------- array/slice start ---------------------
// 迭代遍历数组
func traverseArrayLoop(arr []int) {
	for i := 0; i < len(arr); i++ {
	}
	for i := len(arr) - 1; i >= 0; i-- {
	}
	for index, v := range arr {
		_, _ = index, v
	}
}

// 递归遍历数组
func traverseArrayRecursive(arr []int, index int) {
	if index == len(arr) { // 递归出口
		return
	}
	// 前序位置
	traverseArrayRecursive(arr, index+1)
	// 后序位置
}

//------------- array/slice end -----------------------

//------------- linked list start ---------------------
type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代遍历单向链表
func traverseLinkedListLoop(head *ListNode) {
	for p := head; p != nil; p = p.Next {
	}
}

// 递归遍历单向链表
func traverseLinkedListRecursive(head *ListNode) {
	if head == nil { // 递归出口
		return
	}
	// 前序位置
	traverseLinkedListRecursive(head.Next)
	// 后序位置
}

//------------- linked list end -----------------------
