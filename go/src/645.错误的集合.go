/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	n := len(nums)
	sum, actualSum := 0, 0
	exist := make([]bool, n + 1)
	for _, num := range nums {
		if !exist[num] {
			exist[num] = true
			actualSum += num
		}
		sum += num
	}

	return []int{sum - actualSum, (n + 1) * n / 2 - actualSum}
}
// @lc code=end

