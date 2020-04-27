/*
 * @lc app=leetcode.cn id=546 lang=golang
 *
 * [546] 移除盒子
 */

// @lc code=start
func removeBoxes(boxes []int) int {
	n := len(boxes)

	// dp[l][r][k]，l 表示起始下标，r 表示结束下标，k 表示与第 r 个元素相似的元素（坐标>r）个数，
	// 可以在最后一起合并并得到最终的分数存入 dp[l][r][k] 中
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	return cal(dp, boxes, 0, n - 1, 0)
}

func cal(dp [][][]int, boxes []int, l, r, k int) int {
	if l > r {
		return 0
	}

	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}

	for r > l && boxes[r] == boxes[r - 1] {
		r--
		k++
	}

	val := cal(dp, boxes, l, r - 1, 0) + (k + 1) * (k + 1)
	for i := l; i < r; i++ {
		if boxes[i] == boxes[r] {
			val = max(val, cal(dp, boxes, i + 1, r - 1, 0) + cal(dp, boxes, l, i, k + 1))
		}
	}
	dp[l][r][k] = val
	return dp[l][r][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

