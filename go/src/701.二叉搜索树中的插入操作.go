/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
 func insertIntoBST(root *TreeNode, val int) *TreeNode {
  if root == nil {
    return &TreeNode{Val: val}
  }
  
  cur := root
  for {
    if cur.Val > val {
      if cur.Left == nil {
        cur.Left = &TreeNode{Val: val}
        return root
      } else {
        cur = cur.Left
      }
    } else {
      if cur.Right == nil {
        cur.Right = &TreeNode{Val: val}
        return root
      } else {
        cur = cur.Right
      }
    }
  }
  return root
}
// @lc code=end

