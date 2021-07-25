/*
 * @lc app=leetcode.cn id=1877 lang=golang
 *
 * [1877] 数组中最大数对和的最小值
 */

// @lc code=start
func minPairSum(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	res := nums[0] + nums[n-1]
	i, j := 1, n-2
	for i < j {
		res = max(res, nums[i]+nums[j])
		i++
		j--
	}

	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}

// @lc code=end

