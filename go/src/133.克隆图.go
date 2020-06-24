/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
  if node == nil {
		return nil
	}

	m := map[int]*Node{}
	return dfs(node, m)
}

func dfs(node *Node, m map[int]*Node) *Node {
	root := &Node{ Val: node.Val }
	m[node.Val] = root
	for _, n := range node.Neighbors {
		if _, ok := m[n.Val]; !ok {
			m[n.Val] = dfs(n, m)
		}

		root.Neighbors = append(root.Neighbors, m[n.Val])
	}

	return root
}
// @lc code=end

