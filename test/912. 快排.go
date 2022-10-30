package main

import "fmt"

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
	for i := start + 1; i <= end; i++ {
		if nums[i] < pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[start], nums[index-1] = nums[index-1], nums[start]
	return index - 1
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	sortArray(nums)
	fmt.Println(nums)
}
