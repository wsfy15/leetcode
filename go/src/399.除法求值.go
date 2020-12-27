/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	index := 0
	symbol := map[string]int{}
	for _, equation := range equations {
		for _, str := range equation {
			if _, ok := symbol[str]; !ok {
				symbol[str] = index
				index++
			}
		}
	}

	g := make([][]float64, index)
	for i := 0; i < index; i++ {
		g[i] = make([]float64, index)
	}

	for i := 0; i < len(equations); i++ {
		start, end := symbol[equations[i][0]], symbol[equations[i][1]]
		g[start][end] = values[i]
		g[end][start] = 1 / values[i]
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := symbol[query[0]]; !ok {
			res[i] = -1
			continue
		}
		if _, ok := symbol[query[1]]; !ok {
			res[i] = -1
			continue
		}

		visited := make([]bool, index)
		res[i], _ = doQuery(g, visited, symbol[query[0]], symbol[query[1]])
	}

	return res
}

func doQuery(g [][]float64, visited []bool, start, end int) (float64, bool) {
	if g[start][end] != 0 {
		return g[start][end], true
	}

	visited[start] = true
	for i := 0; i < len(g); i++ {
		if visited[i] || g[start][i] == 0 {
			continue
		}

		if num, ok := doQuery(g, visited, i, end); ok {
			return g[start][i] * num, true
		}
	}

	return -1, false
}

// @lc code=end

