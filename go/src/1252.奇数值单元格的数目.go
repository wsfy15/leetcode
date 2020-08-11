/*
 * @lc app=leetcode.cn id=1252 lang=golang
 *
 * [1252] 奇数值单元格的数目
 */

// @lc code=start
func oddCells(n int, m int, indices [][]int) int {
	row := make([]bool, n)
	col := make([]bool, m)

	for _, index := range indices {
		row[index[0]] = !row[index[0]]
		col[index[1]] = !col[index[1]]
	}

	rr, cc := 0, 0
	for _, val := range row {
		if val {
			rr++
		}
	}
	for _, val := range col {
		if val {
			cc++
		}
	}

	return rr*m + cc*n - rr*cc*2
}

// @lc code=end

