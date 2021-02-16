/*
 * @lc app=leetcode.cn id=1579 lang=golang
 *
 * [1579] 保证图可完全遍历
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

func maxNumEdgesToRemove(n int, edges [][]int) int {
	uf := NewUF(n + 1)
	node, used := n, 0
	for _, edge := range edges {
		if edge[0] == 3 {
			if merge(edge[1], edge[2], uf) {
				node--
				used++
			}
		}
	}

	alice := make([]int, n + 1)
	copy(alice, uf)
	if v, ok := util(node, 1, alice, edges); ok {
		used += v
	} else {
		return -1
	}
	if v, ok := util(node, 2, uf, edges); ok {
		used += v
	} else {
		return -1
	}

	return len(edges) - used
}

func util(n, typ int, uf []int, edges [][]int) (int, bool) {
	used := 0
	for _, edge := range edges {
		if edge[0] == typ {
			if merge(edge[1], edge[2], uf) {
				used++
				n--
			}
		}
	}

	if n == 1 {
		return used, true
	} 
	return 0, false
}
// @lc code=end

