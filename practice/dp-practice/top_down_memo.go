package dppractice

// 自顶向下备忘录法
// 画出递归树
// 可以看到如果没有备忘录，很多节点的值需要重复计算，且n越大，重复计算的成本越高
// 一方面，大量重复计算需要消耗大量CPU时间，算法复杂度高
// 另外一方面，递归存在调用栈，递归越深，内存消耗越大，空间复杂度越高
// 递归算法的时间复杂度：递归树节点总数 * 每次计算的时间复杂度
//
// 采用备忘录法之后，每次计算出一个节点的值，九将其存入备忘录中，下次需要重复计算时，
// 直接从备忘录中取出该值即可，避免了重复计算
// 此时，递归树将变为一颗单边树——只有一个分支的树，时间和空间复杂度都退化为O(n)
func fibonacci(n int) int {
	memo := make([]int, n+1)
	return fib(n, memo)
}

func fib(n int, memo []int) int {
	if n <= 2 {
		memo[n] = 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}
