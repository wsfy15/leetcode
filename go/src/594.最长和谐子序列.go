/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	sort.Ints(nums)
	res, i, j := 0, 0, 0
	for j < n {
		if nums[j] - nums[i] == 1 {
			res = max(res, j - i + 1)
			j++
		} else if nums[j] - nums[i] > 1 {
			i++
		} else {
			j++
		}
	}

	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
// @lc code=end

