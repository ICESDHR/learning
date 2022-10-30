package main

func maxEnvelopes(envelopes [][]int) int {
	nums := sort(envelopes)
	return lengthOfLIS(nums)
}

func sort(envelopes [][]int) []int {
	for i := 0; i < len(envelopes); i++ {
		for j := i + 1; j < len(envelopes); j++ {
			if envelopes[i][0] > envelopes[j][0] {
				envelopes[i], envelopes[j] = envelopes[j], envelopes[i]
			}
			if envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1] {
				envelopes[i], envelopes[j] = envelopes[j], envelopes[i]
			}
		}
	}

	res := []int{}
	for i := 0; i < len(envelopes); i++ {
		res = append(res, envelopes[i][1])
	}
	return res
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < len(nums); i++ {
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
