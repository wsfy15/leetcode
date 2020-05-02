/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	var res [][]string
	dfs(n, 0, 0, 0, 0, []string{}, &res)

	return res
}

func dfs(n, depth int, col, pie, na int, cur []string, res *[][]string) {
	if n == depth {
		dst := make([]string, n)
		copy(dst, cur)
		*res = append(*res, dst)
		return 
	}

	// 取得可用的位
	// bits := (col | pie | na) ^ ((1 << n) - 1) // 不能这么写，因为pie每次左移，可能使1的位置超出n
	bits := ^(col | pie | na) & ((1 << n) - 1) 
	for bits > 0 {
		num := bits & -bits // 取得最低位的1
		sb := strings.Builder{}
		tmp, count := num >> 1, 1
		for tmp > 0 {
			sb.WriteByte('.')
			tmp >>= 1
			count++
		}
		sb.WriteByte('Q')
		for count < n {
			count++
			sb.WriteByte('.')
		}

		dfs(n, depth + 1, col | num, (pie | num) << 1, (na | num) >> 1, append(cur, sb.String()), res)
		bits &= bits - 1 // 清除最低位的1
	}
}
// @lc code=end

