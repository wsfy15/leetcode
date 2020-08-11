/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	counter := map[*TreeNode]int{}
	return util(root, counter)
}

func util(root *TreeNode, counter map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if v, ok := counter[root]; ok {
		return v
	}

	res1 := root.Val
	if root.Left != nil {
		res1 += util(root.Left.Left, counter) + util(root.Left.Right, counter)
	}
	if root.Right != nil {
		res1 += util(root.Right.Left, counter) + util(root.Right.Right, counter)
	}

	res2 := util(root.Left, counter) + util(root.Right, counter)
	counter[root] = max(res1, res2)
	return counter[root]
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

