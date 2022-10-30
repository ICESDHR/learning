package main

import "math"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		dp[i] = max(nums[i-1], dp[i-1]+nums[i-1])
	}

	res := math.MinInt32
	for i := 1; i <= len(nums); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
