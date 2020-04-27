/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
    if len(nums) < 2 {
			return nums
		}
		quicksort(nums, 0, len(nums) - 1)
		return nums
}

func quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(nums, start, end)
	quicksort(nums, start, pivot - 1)
	quicksort(nums, pivot + 1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	i := start
	for j := start; j <= end; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i] 
	return i
}
// @lc code=end

