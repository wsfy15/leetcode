/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	sum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	if n == 0 {
		return NumMatrix{}
	}

	m := len(matrix[0])
	data := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		data[i] = make([]int, m + 1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			data[i][j] = data[i][j - 1] + data[i - 1][j] - data[i - 1][j - 1] + matrix[i - 1][j - 1]
		}
	}
	return NumMatrix{ sum: data }
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := this.sum[row2 + 1][col2 + 1]
	if row1 > 0 {
		sum -= this.sum[row1][col2 + 1]
	}
	if col1 > 0 {
		sum -= this.sum[row2 + 1][col1]
	}
	if row1 > 0 && col1 > 0 {
		sum += this.sum[row1][col1]
	}
	return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

