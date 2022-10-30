package main

func generateMatrix(n int) [][]int {
	left, right, top, bottom := 0, n-1, 0, n-1
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	num := 1
	tar := n * n
	for num <= tar {
		for i := left; i <= right; i++ {
			mat[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			mat[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			mat[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			mat[i][left] = num
			num++
		}
		left++
	}
	return mat
}
