/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

// @lc code=start
var res int

func reversePairs(nums []int) int {
	res = 0
	mergeSort(nums, 0, len(nums)-1)
	return res
}

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)>>1
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	merge(nums, start, mid, end)
}

func merge(nums []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	index := 0
	i, j := start, mid+1
	for i <= mid {
		for j <= end && nums[i] > nums[j]<<1 {
			j++
		}

		res += j - mid - 1
		i++
	}

	i, j = start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp[index] = nums[i]
			i++
		} else {
			tmp[index] = nums[j]
			j++
		}
		index++
	}

	for i <= mid {
		tmp[index] = nums[i]
		index++
		i++
	}

	for j <= end {
		tmp[index] = nums[j]
		index++
		j++
	}

	for i := end; i >= start; i-- {
		index--
		nums[i] = tmp[index]
	}
}

// @lc code=end

