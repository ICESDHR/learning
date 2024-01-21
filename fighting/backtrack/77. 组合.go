package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	if k <= 0 || n <= 0 {
		return res
	}
	track := []int{}
	backtrack(n, k, 1, track, &res)
	return res
}

func backtrack(n, k, start int, track []int, res *[][]int) {
	if k == len(track) {
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}
