/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var count = make([]int, n + 1)
	count[1] = 1
	count[2] = 2

	return util(n, count)
}

func util(n int , count []int) int {
	if count[n] != 0 {
		return count[n]
	}
	
	count[n] = util(n - 1, count) + util(n - 2, count)
	return count[n] 
}
// @lc code=end

