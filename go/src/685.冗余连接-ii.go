/*
 * @lc app=leetcode.cn id=685 lang=golang
 *
 * [685] 冗余连接 II
 */

// @lc code=start
func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges) // n条边 n个节点
	inDegree := make([]int, n+1)
	for _, edge := range edges {
		inDegree[edge[1]]++
	}

	// 最多有一个节点入度为2，那么要删除的边就是这两条边其中一条
	// 如果入度都是1，说明首尾相连了

	for i := n - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 && helper(edges, i) {
			return edges[i]
		}
	}

	for i := n - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 1 && helper(edges, i) {
			return edges[i]
		}
	}

	return nil
}

// except 当前判断的边的编号
// 判断去除edges[except]后，是否还连通
func helper(edges [][]int, except int) bool {
	n := len(edges)
	uf := make([]int, n+1) // 记录每个节点的父节点

	for i := 1; i <= n; i++ {
		uf[i] = i
	}

	count := n
	for i := 0; i < n; i++ {
		if i == except {
			continue
		}

		fx, fy := find(uf, edges[i][0]), find(uf, edges[i][1])
		if fx != fy { // 判断是否属于同一连通分量
			count--
			uf[fy] = fx
		}
	}

	return count == 1
}

func find(uf []int, x int) int {
	if uf[x] == x {
		return x
	}
	uf[x] = find(uf, uf[x])
	return uf[x]
}

// @lc code=end

