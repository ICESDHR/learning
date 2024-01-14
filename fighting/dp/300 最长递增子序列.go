package main

func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j < i; j++ {
			if dp[j-1] > dp[i-1] {
				dp[i] = max(dp[i], dp[j-1]+1)
			}
		}
	}
	return dp[len(nums)]
}
