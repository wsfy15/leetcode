/*
 * @lc app=leetcode.cn id=1631 lang=golang
 *
 * [1631] 最小体力消耗路径
 */

// @lc code=start
var dx []int = []int{0, 0, 1, -1}
var dy []int = []int{1, -1, 0, 0}

type pair struct {
	x int
	y int
}

// 二分 + dfs/bfs
func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	l, r := 0, int(1e6) 
	for l < r {
		mid := l + (r - l) >> 1
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}

		visited[0][0] = true
		// if dfs(0, 0, mid, heights, visited) {
		// 	r = mid
		// } else {
		// 	l = mid + 1
		// }

		if bfs(mid, heights, visited) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func bfs(limit int, heights [][]int, visited [][]bool) bool {
	queue := []pair{pair{0, 0}}
	for len(queue) > 0 {
		node := queue[0]
		if node.x == len(heights) - 1 && node.y == len(heights[0]) - 1 {
			return true
		}
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			x, y := node.x + dx[i], node.y + dy[i]
			if 0 <= x && x < len(heights) && 0 <= y && y < len(heights[0]) && !visited[x][y] && abs(heights[x][y] - heights[node.x][node.y]) <= limit {
				visited[x][y] = true
				queue = append(queue, pair{ x, y})
			}
		}
	}
	
	return false
}

func dfs(curX, curY, limit int, heights [][]int, visited [][]bool) bool {
	if curX == len(heights) - 1 && curY == len(heights[0]) - 1 {
		return true
	}

	for i := 0; i < 4; i++ {
		x, y := curX + dx[i], curY + dy[i]
		if 0 <= x && x < len(heights) && 0 <= y && y < len(heights[0]) && !visited[x][y] && abs(heights[x][y] - heights[curX][curY]) <= limit {
			visited[x][y] = true
			if dfs(x, y, limit, heights, visited) {
				return true
			}
		}
	}

	return false
}

// 并查集 以边为单位，直到左上角与右下角的点在同一个集合
func NewUF(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(node int, uf []int) int {
	if node != uf[node] {
		uf[node] = find(uf[node], uf)
	}
	return uf[node]
}

func merge(a, b int, uf []int) bool {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
	return aroot != broot
}

type edge struct {
	x int
	y int
	weight int
}

func getNodeId(x, y, col int) int {
	return x * col + y
}

func minimumEffortPath3(heights [][]int) int {
	edges := []edge{}
	n, m := len(heights), len(heights[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 {
				edges = append(edges, edge{getNodeId(i, j, m), getNodeId(i - 1, j, m), abs(heights[i][j] - heights[i - 1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{getNodeId(i, j, m), getNodeId(i, j - 1, m), abs(heights[i][j] - heights[i][j - 1])})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	res, index := 0, 0
	uf := NewUF(n * m)
	for find(0, uf) != find(getNodeId(n - 1, m - 1, m), uf) {
		x, y, weight := edges[index].x, edges[index].y, edges[index].weight
		index++

		if merge(x, y, uf) {
			res = max(res, weight)
		}
	}

	return res
}


// dijkstra
type point struct {
	x int
	y int
	diff int
}

type hp struct {
	points []point
	count int
}

func NewHP() hp {
	return hp{
		points: make([]point, 1),
	}
}

func (h *hp) Insert(p point) {
	h.points = append(h.points, p)
	h.count++
	pos := h.count
	for pos / 2 > 0 && h.points[pos / 2].diff > h.points[pos].diff {
		h.points[pos], h.points[pos / 2] = h.points[pos / 2], h.points[pos]
		pos /= 2
	}
}

func (h *hp) Pop() point {
	p := h.points[1]
	h.points[1] = h.points[h.count]
	h.points = h.points[:h.count]
	h.count--

	pos := 1
	for {
		tmp := pos
		if pos * 2 <= h.count && h.points[pos * 2].diff < h.points[pos].diff {
			tmp = tmp * 2
		}
		if pos * 2 + 1 <= h.count && h.points[pos * 2 + 1].diff < h.points[tmp].diff {
			tmp = pos * 2 + 1
		}

		if tmp == pos {
			break
		}
		h.points[tmp], h.points[pos] = h.points[pos], h.points[tmp]
		pos = tmp
	}

	return p
}

func (h hp) Len() int {
	return h.count
}

func minimumEffortPath2(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	diff := make([][]int, n)
	for i := 0; i < n; i++ {
		diff[i] = make([]int, m)
		for j := 0; j < m; j++ {
			diff[i][j] = math.MaxInt32
		}
	}

	diff[0][0] = 0
	h := NewHP()
	h.Insert(point{0, 0, 0})
	for h.Len() > 0 {
		p := h.Pop()
		if p.x == n - 1 && p.y == m - 1 {
			return diff[n - 1][m - 1]
		}

		if diff[p.x][p.y] < p.diff {
			continue
		}

		for i := 0; i < 4; i++ {
			x, y := p.x + dx[i], p.y + dy[i]
			if 0 <= x && x < n && 0 <= y && y < m {
				val := max(diff[p.x][p.y], abs(heights[x][y] - heights[p.x][p.y]))
				if val < diff[x][y] {
					diff[x][y] = val
					h.Insert(point{x, y, val})
				}
			}
		}
	}


	return diff[n - 1][m - 1]
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

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}
// @lc code=end

