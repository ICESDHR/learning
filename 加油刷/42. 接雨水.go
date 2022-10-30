package main

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	res := 0
	left := make([]int, len(height))
	right := make([]int, len(height))

	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > left[i-1] {
			left[i] = height[i]
		} else {
			left[i] = left[i-1]
		}
	}

	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > right[i+1] {
			right[i] = height[i]
		} else {
			right[i] = right[i+1]
		}
	}

	for i := 0; i < len(height); i++ {
		res += min(left[i], right[i]) - height[i]
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
