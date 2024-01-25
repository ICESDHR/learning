package main

func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	result := 0
	for i := 0; i <= len(nums); i++ {
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}
