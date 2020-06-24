/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	n := len(T)
	if n <= 0 {
		return nil
	}

	res := make([]int, n)
	stack := []int{0}
	for i := 1; i < n; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack) - 1]] {
			res[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	for _, v := range stack {
		res[v] = 0
	}

	return res
}
// @lc code=end

