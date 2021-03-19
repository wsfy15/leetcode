/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n < 2 {
		return n
	}

	// 先对width升序排序，width相同对height降序排序，这样在判断的时候只需要考虑height
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || 
		(envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	return max(dp...)
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

