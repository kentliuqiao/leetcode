package dppractice

import "math"

// Problem
/**
假定我们知道Serling公司出售一段长度为 i 英寸的钢条的价格为 Pi（i=1，2，3...）。钢条的长度均为整英寸。
一个价格表样例
---------------------------------------------
长度i   |  1  2  3  4  5  6  7  8  9  10
---------------------------------------------
价格Pi  |  1  5  8  9  10 17 17 20 24 30
---------------------------------------------
钢条切割问题：给定一段长度为 n 英寸的钢条和一个价格表Pi（i=1，2...n），求切割钢条方案，是的销售收益 Rn 最大。
注意，最优解可能不需要切割。
分析：将钢条从左边切割下长度为 i 的一段，只对右边剩下的长度为 n-i 的一段继续切割（递归求解），对左边的一段则不再切割。
即问题的分解方式为，将长度为 n 的一段钢条分解为左边一段，以及剩余部分继续分解的结果。于是可以得到公式：
Rn = max{Pi, Rn-i}(1 <= i <= n)

参考：https://blog.csdn.net/u013309870/article/details/75193592
*/

// 价格表
var p = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

// cutRecursive 递归求解
// p 价格表, n 长度
//
// 思考：画出递归树之后就会发现，递归存在着大量的冗余计算，空间时间复杂度极高
func cutRecursive(p []int, n int) int {
	if n == 0 {
		return 0
	}
	r := math.MinInt64
	for i := 1; i <= n; i++ {
		r = max(r, p[i]+cutRecursive(p, n-i))
	}
	return r
}

// 自顶向下备忘录法
func cutTopDownMemo(p []int, n int) int {
	memo := make([]int, n+1)
	return cutTopDown(p, memo, n)
}

func cutTopDown(p, memo []int, n int) int {
	if n <= 0 {
		return 0
	}
	if memo[n] > 0 {
		return memo[n]
	}
	r := math.MinInt32
	for i := 1; i <= n; i++ {
		r = max(r, p[i]+cutTopDown(p, memo, n-i))
	}
	memo[n] = r
	return r
}

// cutBottomUp
func cutBottomUp(p []int, n int) int {
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		q := -1
		for j := 1; j <= i; j++ {
			q = max(q, p[j]+r[i-j])
		}
		r[i] = q
	}
	return r[n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
