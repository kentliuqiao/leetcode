package dppractice

// bottomUpMemoFibonacci 自底向上备忘录法
// 思考：自顶向下的计算过程，最终还是先计算f(1),f(2)..., 直到f(n-1)和f(n-2)计算完成之后，才能得到f(n)的值
// 那么为何不能直接从f(1),f(2)计算呢，这样岂不是不需要从f(n)到f(1)的调用栈了吗！
// 因此自底向上的备忘录法减少了自顶向下的调用栈的空间消耗，但是其时间复杂度并未减少
func bottomUpMemoFibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	memo := make([]int, n+1)
	memo[1] = 1
	memo[2] = 1
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

// bottomUpNoMemoFibonacci 自底向上状态压缩法
// 思考：f(n) = f(n-1) + f(n-2)，每次计算，只需要之前前两位的值即可，不需要记住之前所有的值如f(n-3)...f(1)
// 因此状态压缩方法进一步减少了计算的空间复杂度，使其减少为O(1)！
func bottomUpNoMemoFibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	n1, n2 := 1, 1
	for i := 3; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n1
}
