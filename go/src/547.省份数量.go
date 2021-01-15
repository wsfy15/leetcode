/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func initUF(n int) []int {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func find(cur int, uf []int) int {
	tmp := cur
	for uf[cur] != cur {
		cur = uf[cur]
	}

	// 将从当前节点到根节点的每一个节点的uf设为根节点 压缩路径
	for tmp != uf[tmp] {
		tmp, uf[tmp] = uf[tmp], cur
	}
	return cur
}

func merge(a, b int, uf []int) bool {
	aroot, broot := find(a, uf), find(b, uf)
	uf[aroot] = broot
	return aroot != broot
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n < 2 {
		return n
	}

	res := n
	uf := initUF(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 && merge(i, j, uf) {
				res--
			}
		}
	}

	return res
}

// @lc code=end

