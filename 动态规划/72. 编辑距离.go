package main

// 暴力，会超时
func minDistance1(word1 string, word2 string) int {
	var dp func(len1, len2 int) int
	dp = func(len1, len2 int) int {
		if len1 == -1 {
			return len2 + 1
		}
		if len2 == -1 {
			return len1 + 1
		}

		if word1[len1] == word2[len2] {
			return dp(len1-1, len2-1)
		} else {
			return min(dp(len1-1, len2)+1,
				dp(len1, len2-1)+1,
				dp(len1-1, len2-1)+1)
		}

	}
	return dp(len(word1)-1, len(word2)-1)
}

// 动规
func minDistance(word1 string, word2 string) int {
	grid := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		grid[i] = make([]int, len(word2))
	}

	var dp func(len1, len2 int) int
	dp = func(len1, len2 int) int {
		if len1 == -1 {
			return len2 + 1
		}
		if len2 == -1 {
			return len1 + 1
		}
		if grid[len1][len2] != 0 {
			return grid[len1][len2]
		}
		if word1[len1] == word2[len2] {
			grid[len1][len2] = dp(len1-1, len2-1)
			return grid[len1][len2]
		} else {
			grid[len1][len2] = min(dp(len1-1, len2)+1,
				dp(len1, len2-1)+1,
				dp(len1-1, len2-1)+1)
			return grid[len1][len2]
		}
	}
	return dp(len(word1)-1, len(word2)-1)
}

func min(a, b, c int) int {
	if a <= b {
		if a <= c {
			return a
		} else {
			return c
		}
	}
	if b >= c {
		return c
	}
	return b
}
