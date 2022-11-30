package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	visited := make(map[int]bool)
	var backtrack func(nums, track []int)
	backtrack = func(nums, track []int) {
		if len(track) == len(nums) {
			res = append(res, append([]int{}, track...))
			return
		}

		for _, num := range nums {
			if visited[num] {
				continue
			}
			visited[num] = true
			track = append(track, num)
			backtrack(nums, track)
			track = track[:len(track)-1]
			visited[num] = false
		}
	}

	backtrack(nums, track)
	return res
}

func main() {
	a := permute([]int{1, 2, 3})
	fmt.Println(a)
}
