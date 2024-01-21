package main

func spiralOrder(matrix [][]int) []int {
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := []int{}
	num := 0
	tar := len(matrix) * len(matrix[0])
	for num < tar {
		for i := left; i <= right && num < tar; i++ {
			num++
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bottom && num < tar; i++ {
			num++
			res = append(res, matrix[i][right])
		}
		right--

		for i := right; i >= left && num < tar; i-- {
			num++
			res = append(res, matrix[bottom][i])
		}
		bottom--

		for i := bottom; i >= top && num < tar; i-- {
			num++
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
