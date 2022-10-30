package main

import "math"

func superEggDrop(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}

	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		dp[1][i] = i
	}
	for i := 1; i <= K; i++ {
		dp[i][1] = 1
	}
	for i := 2; i <= K; i++ {
		for j := 2; j <= N; j++ {
			for x := 1; x <= j; x++ {
				if dp[i][j] == 0 {
					dp[i][j] = math.MaxInt32
				}
				dp[i][j] = min(dp[i][j], max(dp[i][j-x], dp[i-1][x-1])+1)
			}
		}
	}
	return dp[K][N]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
