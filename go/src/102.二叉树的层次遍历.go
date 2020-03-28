/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
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
 func levelOrder(root *TreeNode) [][]int {
  if root == nil {
    return nil
  }
  
  queue := []*TreeNode{root}
  backup := []*TreeNode{}
  var res [][]int
  i := 0
  for len(queue) > 0 {
    cur := queue[0]
    queue = queue[1:]
    
    if len(res) <= i {
      res = append(res, []int{})
    }
    res[i] = append(res[i], cur.Val)
    
    if cur.Left != nil {
      backup = append(backup, cur.Left)
    }
    if cur.Right != nil {
      backup = append(backup, cur.Right)
    }
    
    if len(queue) == 0 {
      queue, backup = backup, queue
      i++
    }
  }
  
  return res
}
// @lc code=end

