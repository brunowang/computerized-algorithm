package dp

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i < coin {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
