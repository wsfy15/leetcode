/*
 * @lc app=leetcode.cn id=1584 lang=golang
 *
 * [1584] 连接所有点的最小费用
 */

// @lc code=start
type Edge struct {
	x    int
	y    int
	diff int
}

type Edges []Edge

func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Less(i, j int) bool {
	return e[i].diff < e[j].diff
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func NewUF(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(n int, uf []int) int {
	if uf[n] != n {
		uf[n] = find(uf[n], uf)
	}
	return uf[n]
}

func merge(a, b int, uf []int) bool {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
	return aroot != broot
}

// Kruskal 针对边  Prim 针对顶点
func minCostConnectPoints(points [][]int) int {
	var edges Edges
	n := len(points)
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			diff := abs(x-points[j][0]) + abs(y-points[j][1])
			edges = append(edges, Edge{
				x:    i,
				y:    j,
				diff: diff,
			})
		}
	}
	sort.Sort(edges)

	uf := NewUF(n)
	res := 0
	for _, edge := range edges {
		if merge(edge.x, edge.y, uf) {
			res += edge.diff
		}
	}

	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}

// @lc code=end

