/*
 * @lc app=leetcode.cn id=803 lang=golang
 *
 * [803] 打砖块
 */

// @lc code=start
type UF struct {
	parent []int
	size   []int // 以该节点为根的子树的节点数量
}

func NewUf(n int) UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return UF{
		parent: parent,
		size:   size,
	}
}

func (this UF) find(n int) int {
	if this.parent[n] != n {
		this.parent[n] = this.find(this.parent[n])
	}
	return this.parent[n]
}

func (this UF) union(a, b int) {
	aroot, broot := this.find(a), this.find(b)
	if aroot != broot {
		this.parent[aroot] = broot
		this.size[broot] += this.size[aroot]
	}
}

func (this UF) getSize(n int) int {
	return this.size[this.find(n)]
}

func hitBricks(grid [][]int, hits [][]int) []int {
	row, col := len(grid), len(grid[0])
	tmp := make([][]int, row) // 将所有击打位置设为0
	for i := 0; i < row; i++ {
		tmp[i] = make([]int, col)
		copy(tmp[i], grid[i])
	}

	for _, hit := range hits {
		x, y := hit[0], hit[1]
		tmp[x][y] = 0
	}

	size := row * col
	uf := NewUf(size + 1) // 最后一个元素为屋顶
	for j := 0; j < col; j++ {
		if tmp[0][j] == 1 {
			uf.union(j, size)
		}
	}

	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if tmp[i][j] == 0 {
				continue
			}

			if tmp[i-1][j] == 1 {
				uf.union(getIndex(i, j, col), getIndex(i-1, j, col))
			}
			if j > 0 && tmp[i][j-1] == 1 {
				uf.union(getIndex(i, j, col), getIndex(i, j-1, col))
			}
		}
	}

	// 根据hits逆序还原
	// 还原后的与屋顶节点相连的节点数 - 还原前的 - 1 即掉落砖块数目
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}
	res := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		if grid[x][y] == 0 {
			continue
		}

		origin := uf.getSize(size) // 还原前与屋顶相连的节点数

		if x == 0 { // 如果是顶层，需要与屋顶相连
			uf.union(y, size)
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if isValid(nx, ny, row, col) && tmp[nx][ny] == 1 {
				uf.union(getIndex(x, y, col), getIndex(nx, ny, col))
			}
		}

		cur := uf.getSize(size)
		res[i] = max(0, cur-origin-1)
		tmp[x][y] = 1
	}

	return res
}

func getIndex(x, y, col int) int {
	return x*col + y
}

func isValid(x, y, row, col int) bool {
	return 0 <= x && x < row && 0 <= y && y < col
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

