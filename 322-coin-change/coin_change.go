package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	return topDownMemo(memo, coins, amount)
}

func topDownMemo(memo, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] > 0 {
		return memo[amount]
	}
	res := math.MaxInt32
	for _, coin := range coins {
		sub := topDownMemo(memo, coins, amount-coin)
		if sub < 0 {
			continue // no solution with this kind of denomination
		}
		res = min(res, 1+sub)
	}
	if res != math.MaxInt32 {
		memo[amount] = res
		return res
	}
	return -1
}

func bottomUp(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dpTable := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		res := math.MaxInt32
		for _, coin := range coins {
			if i-coin < 0 || dpTable[i-coin] < 0 {
				continue
			}
			res = min(res, 1+dpTable[i-coin])
		}
		if res != math.MaxInt32 {
			dpTable[i] = res
		} else {
			dpTable[i] = -1
		}
	}
	return dpTable[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
