/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
var count int
func totalNQueens(n int) int {
	count = 0
	dfs(n, 0, 0, 0, 0)
	return count
}

// row：当前处理到的行
// col：位X表示第X列上是否有皇后
// pie：从撇的角度记录当前行哪些列上不能使用，体现在进入到下一行时，对pie左移一位
// na：从捺的角度记录当前行哪些列上不能使用，体现在进入到下一行时，对na右移一位
func dfs(n, row, col, pie, na int) {
	if row >= n {
		count++
		return
	}

	// 得到所有可用的位
	// ^(col | pie | na) 将已用的位变为0，其余位变成1.但是从最高位到第n位也全是1
	// (1 << n) - 1 第n-1位到第0位都是1，其他位都是0. &之后清除高位多余的1
	bits := ^(col | pie | na) & ((1 << n) - 1) 
	for bits > 0 {
		// -bits，负数，使用补码表示，即bits所有位取反然后 + 1
		p := bits & -bits // 取最低位的1
		dfs(n, row + 1, col | p, (pie | p) << 1, (na | p) >> 1)
		bits &= bits - 1 // 清除最低位的1
	}
}
// @lc code=end

