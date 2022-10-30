package main

import "sort"

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	mid := len(nums) / 2
	return nums[mid]
}

func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}

// 时间O(N),空间O(1)
