package main

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(s)+1)
	}

	for i := 1; i <= len(s); i++ {
		dp[i][i] = 1
	}

	for i := len(s); i >= 1; i-- {
		for j := i + 1; j <= len(s); j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[1][len(s)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
