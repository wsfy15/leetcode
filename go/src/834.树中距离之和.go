/*
 * @lc app=leetcode.cn id=834 lang=golang
 *
 * [834] 树中距离之和
 */

// @lc code=start
func sumOfDistancesInTree(N int, edges [][]int) []int {
	g := make([][]int, N)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}

	// nodeNum 每个节点所在子树节点数量，包含自己
	nodeNum, distNum := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		nodeNum[i] = 1
	}

	var postOrder func(root, parent int)
	var preOrder func(root, parent int)

	// 先求节点到其子树所有节点的距离
	postOrder = func(root, parent int) {
		for _, node := range g[root] {
			if node == parent {
				continue
			}

			postOrder(node, root)
			nodeNum[root] += nodeNum[node]
			distNum[root] += distNum[node] + nodeNum[node]
		}
	}

	// distNum[0]是正确的，根据其求其他节点的distNum
	preOrder = func(root, parent int) {
		for _, node := range g[root] {
			if node == parent {
				continue
			}

			// 从父节点移动到子节点，访问子节点整个子树的距离减少的值是子树节点个数
			// 访问子节点子树以外的节点，增加的距离是剩余节点数量
			distNum[node] = distNum[root] - nodeNum[node] + N - nodeNum[node]
			preOrder(node, root)
		}
	}

	postOrder(0, -1)
	preOrder(0, -1)

	return distNum
}

// @lc code=end

