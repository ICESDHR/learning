package main

import (
	"fmt"
	"sort"
)

func merge1(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}

// 更简洁的写法
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func main() {
	fmt.Println(merge([]int{}, 0, []int{}, 0))
}
