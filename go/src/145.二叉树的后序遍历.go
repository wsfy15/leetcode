/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
 func postorderTraversal(root *TreeNode) []int {
  if root == nil {
    return nil
  }
  seq := []int{}
  stack := []*TreeNode{root}
  
  for len(stack) > 0 {
    cur := stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]
    seq = append([]int{cur.Val}, seq...) // 加到seq最前面
    
    if cur.Left != nil {
      stack = append(stack, cur.Left)
    }
    if cur.Right != nil {
      stack = append(stack, cur.Right)
    }
    
  }
  
  return seq
}

var res []int
func postorderTraversal2(root *TreeNode) []int {
  if root == nil {
    return nil
  }
  res = []int{}
  postorder(root)
  return res
}

func postorder(root *TreeNode) {
  if root == nil {
    return
  }
  
  postorder(root.Left)
  postorder(root.Right)
  res = append(res, root.Val)
}
// @lc code=end

