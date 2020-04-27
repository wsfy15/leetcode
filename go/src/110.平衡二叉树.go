/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
 func isBalanced(root *TreeNode) bool {
  if root == nil {
    return true
  }
  return height(root) ！= -1
}

func height(root *TreeNode) int {
  if root == nil {
    return 0
  }
  
  left := height(root.Left)
  right := height(root.Right)
  if left == -1 || right == -1 {
    return -1
  }
  
  if left < right {
    if right - left > 1 {
      return -1
    }
    return right + 1
  } 
  if left - right > 1 {
    return -1
  }
  return left + 1
}
// @lc code=end

