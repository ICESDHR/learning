package main

import (
	"fmt"
	"math"
)

func minDistance(word1 string, word2 string) int {
	matrix := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		matrix[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = min(matrix[i-1][j]+1, matrix[i-1][j-1]+1, matrix[i][j-1]+1)
			}
		}
	}
	return matrix[len(word1)][len(word2)]
}

func main() {
	word1 := ""
	word2 := "a"
	fmt.Println(minDistance(word1, word2))
}

func min(nums ...int) int {
	result := math.MaxInt
	for _, v := range nums {
		if v < result {
			result = v
		}
	}
	return result
}
