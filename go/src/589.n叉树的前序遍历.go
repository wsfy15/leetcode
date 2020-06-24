/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		stack= stack[:len(stack) - 1]

		res = append(res, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}
	}

	return res
}

func preorder2(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	preTravel(root, &res)
	return res
}

func preTravel(root *Node, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	for _, child := range root.Children {
		preTravel(child, res)
	}
}
// @lc code=end

