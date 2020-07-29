/*
 * @lc app=leetcode.cn id=849 lang=golang
 *
 * [849] 到最近的人的最大距离
 */

// @lc code=start
func maxDistToClosest(seats []int) int {
	n := len(seats)
	lastOne := 0
	for lastOne < n && seats[lastOne] == 0 {
		lastOne++
	}
	res := lastOne
	for i := lastOne + 1; i < n; i++ {
		if seats[i] == 1 {
			res = max(res, (i - lastOne) / 2)
			lastOne = i
		}
	}
	if seats[n - 1] == 0 {
		res = max(res, n - 1 - lastOne)
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

