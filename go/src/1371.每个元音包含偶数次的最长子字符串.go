/*
 * @lc app=leetcode.cn id=1371 lang=golang
 *
 * [1371] 每个元音包含偶数次的最长子字符串
 */

// @lc code=start
func findTheLongestSubstring(s string) int {
	// 用5个bit表示五个元音字母出现的奇偶，一共有32种情况
	dp := make([]int, 32) // 记录状态i第一次出现时的下标
	dp[0] = -1 // 只有状态0满足5个元音字母出现偶数次
	for i := 1; i < 32; i++ {
		dp[i] = math.MaxInt32
	}
	n := len(s)
	res, state := 0, 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'a':
			state ^= 1
		case 'e':
			state ^= 2
		case 'i':
			state ^= 4
		case 'o':
			state ^= 8
		case 'u':
			state ^= 16
		default:
			break
		}

		if dp[state] == math.MaxInt32 {
			dp[state] = i
		} else {
			// 当s[:i] 与 s[:j] 对应的state相同时，说明s[i+1:j]的状态为0
			res = max(res, i - dp[state])
		}
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

