/*
 * @lc app=leetcode.cn id=960 lang=golang
 *
 * [960] 删列造序 III
 */

// @lc code=start
func minDeletionSize(A []string) int {
	n := len(A)
	chars := len(A[0])
	if chars == 1 {
		return 0
	}

	// 只有每个字符串的 str[i] <= str[j] (i<j)时，inc[i][j] 才为true
	inc := make([][]bool, chars)
	for i := 0; i < chars; i++ {
		inc[i] = make([]bool, chars)
	}

	for i := 0; i < chars; i++ {
		for j := i + 1; j < chars; j++ {
			inc[i][j] = true
			for k := 0; k < n; k++ {
				if A[k][i] > A[k][j] { // 不符合字典序
					inc[i][j] = false
					break
				}
			}
		}
	}


	// 得到最长路径
	max := 0
	dp := make([]int, chars)
	for i := 0; i < chars; i++ { // 以每个索引为起点
		tmp := getMaxLen(inc, dp, i, chars)
		if tmp > max {
			max = tmp
		}
	}

	// max为路径长度，加1后为路径上字符数 
	// 字符数减去最长路径字符数就是删除索引的数量
	return chars - (max + 1)
}

func getMaxLen(inc [][]bool, dp []int, i, chars int) int {
	if dp[i] > 0 {
		return dp[i]
	}

	max := 0
	for j := i + 1; j < chars; j++ {
		if inc[i][j] {
			tmp := 1 + getMaxLen(inc, dp, j, chars)
			if tmp > max {
				max = tmp
			}
		}
	}

	dp[i] = max
	return max
}

// @lc code=end

