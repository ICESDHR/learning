package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := make(map[int]struct{})
	track := []int{}
	backtrack(nums, visited, track, &res)
	return res
}

func backtrack(nums []int, visited map[int]struct{}, track []int, res *[][]int) {
	if len(nums) == len(track) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; ok {
			continue
		}

		visited[nums[i]] = struct{}{}
		track = append(track, nums[i])
		backtrack(nums, visited, track, res)
		track = track[:len(track)-1]
		delete(visited, nums[i])
	}
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
