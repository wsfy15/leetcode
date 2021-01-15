/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start
func NewUf(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(index int, uf []int) int {
	tmp := index
	for uf[index] != index {
		index = uf[index]
	}

	for tmp != uf[tmp] {
		uf[tmp], tmp = index, uf[tmp]
	}
	return index
}

func merge(a, b int, uf []int) {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUf(n + 1)
	for _, edge := range edges {
		if find(edge[0], uf) != find(edge[1], uf) {
			merge(edge[0], edge[1], uf)
		} else {
			return edge
		}
	}

	return nil
}

// @lc code=end

