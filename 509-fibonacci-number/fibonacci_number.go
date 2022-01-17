package fibonaccinumber

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibMemo(n int) int {
	memo := make([]int, n+1)
	return fibMemoDP(memo, n)
}

func fibMemoDP(memo []int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if memo[n] > 0 {
		return memo[n]
	}
	memo[n] = fibMemoDP(memo, n-1) + fibMemoDP(memo, n-2)
	return memo[n]
}

func fibBottomUp(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	x, y := 0, 1
	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
