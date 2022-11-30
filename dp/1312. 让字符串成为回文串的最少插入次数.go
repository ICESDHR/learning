package main

func minInsertions(s string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(s)+1)
	}

	for i := len(s); i >= 1; i-- {
		for j := i + 1; j <= len(s); j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j]+1, dp[i][j-1]+1)
			}
		}
	}

	return dp[1][len(s)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
