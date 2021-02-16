/*
 * @lc app=leetcode.cn id=778 lang=golang
 *
 * [778] 水位上升的泳池中游泳
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

func merge(a, b int, uf []int) {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
}

type edge struct {
	x int
	y int
	weight int
}

func getNodeNumber(x, y, col int) int {
	return x * col + y
}

func swimInWater(grid [][]int) int {
	n := len(grid)

	edges := []edge{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				edges = append(edges, edge{getNodeNumber(i, j, n), getNodeNumber(i - 1, j, n), max(grid[i][j], grid[i - 1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{getNodeNumber(i, j, n), getNodeNumber(i, j - 1, n), max(grid[i][j], grid[i][j - 1])})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	index := 0
	uf := NewUF(n * n)
	for find(0, uf) != find(n * n - 1, uf) {
		merge(edges[index].x, edges[index].y, uf)
		index++
	}
	return edges[index - 1].weight
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

