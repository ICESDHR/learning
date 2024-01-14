package main

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) <= 0 || len(text2) <= 0 {
		return 0
	}

	matrix := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		matrix[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = max(matrix[i-1][j], matrix[i][j-1])
			}
		}
	}
	return matrix[len(text1)][len(text2)]
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
