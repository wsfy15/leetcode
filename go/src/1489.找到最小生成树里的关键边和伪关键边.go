/*
 * @lc app=leetcode.cn id=1489 lang=golang
 *
 * [1489] 找到最小生成树里的关键边和伪关键边
 */

// @lc code=start
func NewUF(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(node int, uf []int) int {
	if uf[node] != node {
		uf[node] = find(uf[node], uf)
	}
	return uf[node]
}

func merge(a, b int, uf []int) bool {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
	return aroot != broot
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)
	for i := 0; i < m; i++ {
		edges[i] = append(edges[i], i)
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	minCost := kruskal1(edges, n, -1)
	res := make([][]int, 2)
	for i := 0; i < m; i++ {
		if kruskal1(edges, n, i) > minCost {
			res[0] = append(res[0], edges[i][3])
		} else if kruskal2(edges, n, i) == minCost {
			res[1] = append(res[1], edges[i][3])
		}
	}

	return res
}

// 计算最小生成树开销，排除node边
func kruskal1(edges [][]int, n, node int) int {
	m, res, count := len(edges), 0, 0
	uf := NewUF(n)
	for i := 0; i < m; i++ {
		if i == node {
			continue
		}

		edge := edges[i]
		x, y := edge[0], edge[1]
		if merge(x, y, uf) {
			count++
			res += edge[2]
		}
	}

	if count == n-1 {
		return res
	}
	return math.MaxInt32
}

// 计算最小生成树开销，必须包含node边
func kruskal2(edges [][]int, n, node int) int {
	m, res, count := len(edges), 0, 0
	uf := NewUF(n)
	for i := 0; i < m; i++ {
		if i == node {
			merge(edges[i][0], edges[i][1], uf)
			res += edges[i][2]
			count++
			break
		}
	}

	for i := 0; i < m; i++ {
		edge := edges[i]
		x, y := edge[0], edge[1]
		if merge(x, y, uf) {
			count++
			res += edge[2]
		}
	}

	if count == n-1 {
		return res
	}
	return math.MaxInt32
}

// @lc code=end

