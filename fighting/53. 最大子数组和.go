package main

import "math"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		dp[i] = nums[i]
	}
	for i := 1; i < len(nums)-1; i++ {
		dp[i] = max(dp[i], dp[i-1]+nums[i])
	}
	result := math.MinInt
	for i := 0; i < len(nums)-1; i++ {
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func max(nums ...int) int {
	maxNum := math.MinInt
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
