/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	states := make([]int, num + 1)
	for i := 1; i <= num; i++ {
		// i & (i - 1)：清除最低位的1
		// 看看清除后还有多少个1，再加上最低位的一个1
		states[i] = states[i & (i - 1)] + 1
	}
	return states
}
// @lc code=end

