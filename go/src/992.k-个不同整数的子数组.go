/*
 * @lc app=leetcode.cn id=992 lang=golang
 *
 * [992] K 个不同整数的子数组
 */

// @lc code=start
func subarraysWithKDistinct(A []int, K int) int {
	return atMostK(A, K) - atMostK(A, K - 1)
}

// 不同整数个数小于等于k的子数组数量
func atMostK(nums []int, k int) int {
	n, res, l, diff := len(nums), 0, 0, 0
	counter := make([]int, n + 1)
	for r := 0; r < n; r++ {
		if counter[nums[r]] == 0 {
			diff++
		}
		counter[nums[r]]++

		for diff > k {
			counter[nums[l]]--
			if counter[nums[l]] == 0 {
				diff--
			}
			l++
		}

		res += r - l + 1
	}

	return res
}

// @lc code=end

