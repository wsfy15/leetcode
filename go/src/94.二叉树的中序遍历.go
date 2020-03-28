/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
 func inorderTraversal(root *TreeNode) []int {
  if root == nil {
    return nil
  }
  seq := []int{}
  stack := []*TreeNode{}
  cur := root
  for len(stack) > 0 || cur != nil {
    for cur != nil {
      stack = append(stack, cur)
      cur = cur.Left
    }
    
    cur = stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]

    seq = append(seq, cur.Val)
    cur = cur.Right
  }
  return seq
}

var res []int
func inorderTraversal2(root *TreeNode) []int {
  if root == nil {
    return nil
  }
  res = []int{}
  inorder(root)
  return res
}

func inorder(root *TreeNode) {
  if root == nil {
    return
  }
  
  inorder(root.Left)
  res = append(res, root.Val)
  inorder(root.Right)
}
// @lc code=end

