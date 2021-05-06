/*
 * @lc app=leetcode.cn id=1238 lang=golang
 *
 * [1238] 循环码排列
 */

// @lc code=start
func circularPermutation(n int, start int) []int {
	if n == 0 {
		return []int{0}
	}

	nums := grayCode(n)
	for i, v := range nums {
		if v == start {
			reverse(nums, 0, i-1)
			reverse(nums, i, len(nums)-1)
			reverse(nums, 0, len(nums)-1)
			break
		}
	}

	return nums
}

func grayCode(n int) []int {
	res := make([]int, 1<<n)
	res[0], res[1] = 0, 1
	for i := 2; i <= n; i++ {
		for j := 1 << (i - 1); j < 1<<i; j++ {
			res[j] = res[1<<i-j-1] + 1<<(i-1)
		}
	}

	return res
}

func reverse(nums []int, start, end int) {
	if start >= end {
		return
	}
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// @lc code=end

