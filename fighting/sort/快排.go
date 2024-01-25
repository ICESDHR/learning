package sort

import "math/rand"

//解题思路
//测试用例更新后增加了例如 5 万个 2 这种极端测试用例，这种特殊用例使得基本的快排算法时间复杂度下降到 O(n^2)。但只需要在原有的代码上需要做一些小的修改就可以应付这种极端用例：
//
//随机选取“支点”来对付基本有序的用例
//频繁交换支点来对付基本一致的用例
//以双指针扫描为例，原来选取第一个值作为支点，现在改为随机选取支点的值。原来扫描时，为了减少交换次数，只对严格大于（或小于）的数据才进行交换，现在改为大于等于（或小于等于）即交换，这样就能保证当数据基本相同时，左右扫描交替进行，将支点最终停留在靠近数组中间的位置。
//
//具体的改动，在代码中注释出来了。

func sortArray(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quickSort(arr, left, partitionIndex-1)
		quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	p := rand.Intn(right-left+1) + left
	arr[left], arr[p] = arr[p], arr[left]
	for left < right {
		for arr[left] < arr[right] && left < right {
			right--
		} // 修改原来的 nums[right] >= nums[left]，增加交换频率
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
		for arr[left] < arr[right] && left < right {
			left++
		} // 修改原来的 nums[right] >= nums[left]，增加交换频率
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}
	return left
}
