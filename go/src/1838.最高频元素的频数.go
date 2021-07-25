/*
 * @lc app=leetcode.cn id=1838 lang=golang
 *
 * [1838] 最高频元素的频数
 */

// @lc code=start
func maxFrequency(nums []int, k int) int {
	n, res, sum := len(nums), 0, 0
	sort.Ints(nums)

	l, r := 0, 0
	for r < n {
		for nums[r]*(r-l)-sum > k {
			sum -= nums[l]
			l++
		}
		sum += nums[r]
		res = max(res, r-l+1)
		r++
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

