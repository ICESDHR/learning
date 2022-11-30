package main

import "fmt"

func threeSum(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		tmp := twoSum(nums[i+1:], 0-nums[i])
		if len(tmp) != 0 {
			for j := 0; j < len(tmp); j++ {
				res = append(res, append(tmp[j], nums[i]))
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	if len(nums) < 2 {
		return nil
	}
	res := [][]int{}
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := tmp[target-nums[i]]; ok {
			res = append(res, []int{target - nums[i], nums[i]})
		}
		tmp[nums[i]] = i
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
