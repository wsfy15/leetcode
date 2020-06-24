/*
 * @lc app=leetcode.cn id=983 lang=golang
 *
 * [983] 最低票价
 */

// @lc code=start
func mincostTickets(days []int, costs []int) int {
	n := len(days)
	nums := make([]int, days[n - 1] + 30)
	j := 0 // days的下标
	for i := 30; i < days[n - 1] + 30; i++ {
		if days[j] != i - 29 {
			nums[i] = nums[i - 1]
		} else {
			nums[i] = min(nums[i - 1] + costs[0], nums[i - 7] + costs[1], nums[i - 30] + costs[2])
			j++
		}
	}

	return nums[days[n - 1] + 29]
}

func min(a ...int) int {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
// @lc code=end

