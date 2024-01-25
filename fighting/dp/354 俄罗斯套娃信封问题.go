package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	nums := []int{}
	for i := 0; i < len(envelopes); i++ {
		nums = append(nums, envelopes[i][1])
	}
	return lengthOfLIS1(nums)
}

func lengthOfLIS1(nums []int) int {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := [][]int{}
	nums = append(nums, []int{4, 5})
	nums = append(nums, []int{4, 6})
	nums = append(nums, []int{6, 7})
	nums = append(nums, []int{2, 3})
	nums = append(nums, []int{1, 1})
	a := maxEnvelopes(nums)
	fmt.Println(a)
}
