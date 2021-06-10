/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n, m := len(matrix), len(matrix[0])
	prefix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prefix[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	res := 0
	for up := 1; up <= n; up++ {
		for bottom := up; bottom <= n; bottom++ {
			// 固定上下边界，移动右边界，并用hash表记录子矩阵和的出现次数
			counter := map[int]int{}
			for r := 1; r <= m; r++ {
				cur := prefix[bottom][r] - prefix[up-1][r]
				if cur == target {
					res++
				}
				if v, ok := counter[cur-target]; ok {
					res += v
				}
				counter[cur]++
			}
		}
	}

	return res
}

func numSubmatrixSumTarget2(matrix [][]int, target int) int {
	n, m := len(matrix), len(matrix[0])
	prefix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prefix[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for p := 0; p < i; p++ { // row
				for q := 0; q < j; q++ { // col
					if prefix[i][j]-prefix[i][q]-prefix[p][j]+prefix[p][q] == target {
						res++
					}
				}
			}
		}
	}
	return res
}

// @lc code=end

