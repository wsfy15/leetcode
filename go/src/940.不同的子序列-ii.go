/*
 * @lc app=leetcode.cn id=940 lang=golang
 *
 * [940] 不同的子序列 II
 */

// @lc code=start
func distinctSubseqII(S string) int {
	n, sum := len(S), 0 // sum记录当前非空子序列数量，相当于endsWith数组和
	var mod int = 1e9 + 7
	endsWith := make([]int, 26) // 记录以每个字符结尾的字符串的个数

	for i := 0; i < n; i++ {
		index := S[i] - 'a'

		// 减去是为了避免重复计算，减去的序列是 加上当前字符后 被包括在 当前以该字符结尾的序列中，这些序列加上当前字符后就会产生重复
		// 加1是指只有当前字符的序列，这个序列没办法通过其他非空序列末尾加上当前字符得到
		added := (sum - endsWith[index] + 1) % mod 
		sum = (sum + added) % mod

		// 新增的序列数量加上当前的序列数量
		endsWith[index] = (endsWith[index] + added) % mod 
	}

	return (sum + mod) % mod
}
// @lc code=end

