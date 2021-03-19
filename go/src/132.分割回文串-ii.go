/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	n := len(s)
	isPali := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPali[i] = make([]bool, n)
		isPali[i][i] = true
	}

	for right := 1; right < n; right++ {
		for left := 0; left < right; left++ {
			if s[left] == s[right] && (right - left == 1 || isPali[left + 1][right - 1]) {
				isPali[left][right] = true
			}
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ { // 初始化分割次数为s[:i] 长度
		dp[i] = i
	}

	for i := 1; i < n; i++ {
		if isPali[0][i] {
			dp[i] = 0
		} else {
			for k := 1; k <= i; k++ {
				if isPali[k][i] {
					dp[i] = min(dp[i], dp[k - 1] + 1)
				}
			}
		}
	}

	return dp[n - 1]
}

// TLE
func minCut2(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 0
	}

	for k := 2; k <= n; k++ {
		for i := 0; i <= n - k; i++ {
			j := i + k - 1
			if j - i == 1 {
				if s[i] == s[j] {
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = math.MaxInt32
				if s[i] == s[j] && dp[i + 1][j - 1] == 0  {
					dp[i][j] = 0
				} else {
					for m := i; m < j; m++ {
						dp[i][j] = min(dp[i][j], dp[i][m] + dp[m + 1][j] + 1)
					}
				}
		}
		}
	}

	return dp[0][n - 1]
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

