/*
 * @lc app=leetcode.cn id=837 lang=golang
 *
 * [837] 新21点
 */

// @lc code=start
// >=K 分时，就停止抽取数字
// 分数不超过 N 的概率是多少？
// 每次从[1, W] 的范围中随机获得一个整数作为分数，每次抽取都是独立的，结果具有相同的概率
func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, N + 1)
	dp[0] = 1
	dp[1] = 1 / float64(W)
	for i := 2; i <= N; i++ {
		dp[i] = dp[i - 1]
		if i <= K {
			dp[i] += dp[i - 1] / float64(W)
		}
		if i - W - 1 >= 0 {
			dp[i] -= dp[i - W - 1] / float64(W)
		}
	}
	return sum(dp[K:N+1]...)
}

func sum(a ...float64) float64 {
	var res float64
	for _, v := range a {
		res += v
	}
	return res
}
// @lc code=end

