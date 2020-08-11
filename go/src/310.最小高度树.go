/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// 从叶子节点开始往上找，一层一层
	// 最后只剩下一个或两个节点，即为根节点
	degree := make([]int, n)
	graph := make([][]int, n)
	for _, edge := range edges {
		degree[edge[0]]++
		degree[edge[1]]++
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	queue := []int{} // 保存度数为1的节点
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	rst := n
	for rst > 2 {
		rst -= len(queue)
		newQueue := []int{}
		for _, node := range queue {
			for _, peer := range graph[node] {
				degree[peer]--
				if degree[peer] == 1 {
					newQueue = append(newQueue, peer)
				}
			}
		}

		queue = newQueue
	}

	return queue
}
// @lc code=end

