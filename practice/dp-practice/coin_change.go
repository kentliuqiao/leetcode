package dppractice

import "math"

/**
凑零钱问题
给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，
问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。算法的函数签名如下：
coins 中是可选硬币面值，amount 是目标金额
int coinChange(int[] coins, int amount);
*/

func coinChange(coins []int, amount int) int {
	return topDown(coins, amount)
}

func topDown(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := topDown(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}
	return res
}

func coinChangeWithMemo(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = math.MinInt32
	}
	return topDownMemo(coins, memo, amount)
}

func topDownMemo(coins, memo []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != math.MinInt32 {
		return memo[amount]
	}
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := topDownMemo(coins, memo, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}
	if res != math.MaxInt32 {
		memo[amount] = res
	} else {
		memo[amount] = -1 // 无解
	}
	return res
}

func coinChangeBottomUp(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	// dpTable 含义：索引 index 表示输入金额amount，dpTable[index]表示输入金额为index时最少需要的硬币数。
	dpTable := make([]int, amount+1)
	for i := 1; i < len(dpTable); i++ {
		res := math.MaxInt32
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			res = min(res, dpTable[i-coin]+1)
		}
		dpTable[i] = res
	}
	if dpTable[amount] > 0 {
		return dpTable[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
