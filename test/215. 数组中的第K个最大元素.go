package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		p := partition(nums, start, end)
		if p > len(nums)-k {
			end = p - 1
		} else if p < len(nums)-k {
			start = p + 1
		} else {
			return nums[p]
		}
	}
	return 0
}
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) []int {
	if start < end {
		p := partition(nums, start, end)
		quickSort(nums, start, p-1)
		quickSort(nums, p+1, end)
	}
	return nums
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	index := start + 1
	for i := index; i <= end; i++ {
		if nums[i] < pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[start], nums[index-1] = nums[index-1], nums[start]
	return index - 1
}

func main() {
	fmt.Println(findKthLargest([]int{1}, 1))
}
