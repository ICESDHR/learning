package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := math.MinInt32
	dp := make([]int, len(nums)+1)
	dp[0] = math.MinInt32
	for i := 1; i <= len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i-1], nums[i-1])
	}

	for i := 0; i <= len(nums); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxSubArray([]int{-2, -1}))
}
