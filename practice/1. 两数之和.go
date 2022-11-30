package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := tmp[target-nums[i]]; ok {
			return []int{tmp[target-nums[i]], i}
		}
		tmp[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
