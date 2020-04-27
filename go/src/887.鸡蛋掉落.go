/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */

// @lc code=start
func superEggDrop(K int, N int) int {
	// dp状态为(K, N)，即(剩下的鸡蛋数，需要考虑的楼层数)
	// 在X楼扔鸡蛋，可能碎，状态变为(K-1, X-1); 没碎的话，状态变为(K, N-X)
	// 由于不知道碎没碎，所以需要取两种情况的最大值
	// 然后遍历所有楼层，得到所有楼层这两种情况最大值 里面的最小值
	// dp[k][n] = 1 + min( 1 <= x <= N max(dp[k-1][x-1], dp[k][n-x])  )
	m := make(map[int]int)
	return dp(K, N, m)
}

func dp(k, n int, m map[int]int) int {
	if k == 1 { // 只有一个鸡蛋，只能逐层试
		return n
	}

	if n <= 1 { // 只有一层楼，只需试一次
		return 1
	}

	if v, ok := m[n*100 + k]; ok {
		return v
	}

	low, high := 1, n
	for low + 1 < high {
		mid := low + (high - low) >> 1
		// dp[k-1][x-1], dp[k][n-x] 分别是对于x（当前楼层）的单调递增和单调递减函数
		// 因此利用二分查找找到两个函数的交点
		f1, f2 := dp(k-1, mid-1, m), dp(k, n-mid, m)
		if f1 < f2 {
			low = mid
		} else if f1 > f2 {
			high = mid
		} else {
			low, high = mid, mid
		}
	}

	m[n*100 + k] = 1 + min(max(dp(k-1, low - 1, m), dp(k, n - low, m)),
		max(dp(k-1, high - 1, m), dp(k, n - high, m)))

	return m[n*100 + k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

