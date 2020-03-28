/*
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] 二叉搜索树中的搜索
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
 func searchBST(root *TreeNode, val int) *TreeNode {
  for root != nil {
    if root.Val == val {
      return root
    }
    
    if root.Val < val {
      root = root.Right
    } else {
      root = root.Left
    }
  }
  
  return root
}
// @lc code=end

