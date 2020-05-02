/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	seq := []int{}
	stack := []*TreeNode{root}
	
	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		seq = append(seq, cur.Val)
		
		// 先入栈右节点，再左节点
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	
	return seq
}

var res []int

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res = []int{}
	preorder(root)
	return res
}
	
func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	
	res = append(res, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

// @lc code=end

