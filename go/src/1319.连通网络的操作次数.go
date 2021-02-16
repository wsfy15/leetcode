/*
 * @lc app=leetcode.cn id=1319 lang=golang
 *
 * [1319] 连通网络的操作次数
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

func makeConnected(n int, connections [][]int) int {
	m := len(connections)
	if m < n-1 {
		return -1
	}

	uf := NewUF(n)
	parts := n
	for _, connection := range connections {
		a, b := connection[0], connection[1]
		if merge(a, b, uf) {
			parts--
		}
	}

	return parts - 1
}

// @lc code=end

