/*
 * @lc app=leetcode.cn id=1042 lang=golang
 *
 * [1042] 不邻接植花
 */

// @lc code=start
func gardenNoAdj(N int, paths [][]int) []int {
	res := make([]int, N)
	g := make([][]int, N + 1)
	for _, path := range paths {
		x, y := path[0], path[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	for i := 1; i <= N; i++ {
		used := [5]bool{}
		for _, j := range g[i] {
			used[res[j-1]] = true
		}
		for j := 1; j < 5; j++ {
			if !used[j] {
				res[i-1] = j
				break
			}
		}
	}

	return res
}
// @lc code=end

