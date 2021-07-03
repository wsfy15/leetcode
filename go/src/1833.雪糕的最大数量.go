/*
 * @lc app=leetcode.cn id=1833 lang=golang
 *
 * [1833] 雪糕的最大数量
 */

// @lc code=start
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	res := 0
	for i := 0; i < len(costs); i++ {
		if costs[i] > coins {
			break
		}
		coins -= costs[i]
		res++
	}

	return res
}

// @lc code=end

