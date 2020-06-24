/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func postorder(root *Node) []int {
	// 根右左的顺序遍历，然后把结果reverse
	if root == nil {
		return nil
	}

	var res []int
	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		stack= stack[:len(stack) - 1]

		res = append(res, cur.Val)
		stack = append(stack, cur.Children...)
	}

	for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func postorder2(root *Node) []int {
	res := []int{}
	postTravel(root, &res)
	return res
}

func postTravel(root *Node, res *[]int) {
	if root == nil {
		return
	}

	for _, child := range root.Children {
		postTravel(child, res)
	}
	*res = append(*res, root.Val)
}
// @lc code=end

