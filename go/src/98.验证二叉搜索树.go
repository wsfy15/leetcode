/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
 func isValidBST(root *TreeNode) bool {
  // 利用中序遍历
  if root == nil {
    return true
  }
  
  inorder := []int{}
  stack := []*TreeNode{}
  cur := root
  for len(stack) > 0 || cur != nil {
    for cur != nil {
      stack = append(stack, cur)
      cur = cur.Left
    }
    
    cur = stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]
    
    inorder = append(inorder, cur.Val)
    cur = cur.Right
  }
  
  for i := 0; i < len(inorder) - 1; i++ {
    if inorder[i] >= inorder[i+1] {
      return false
    }
  }
  return true
}

func isValidBST2(root *TreeNode) bool {
  if root == nil {
    return true
  }
  return isValid(root, math.MaxInt64, math.MinInt64)
  // return isValid(root.Left, root.Val, math.MinInt32) && isValid(root.Right, math.MaxInt32, root.Val)
}

func isValid(cur *TreeNode, max, min int64) bool {
  if cur == nil {
    return true
  }
  
  if int64(cur.Val) < max && int64(cur.Val) > min {
    return isValid(cur.Left, int64(cur.Val), min) && isValid(cur.Right, max, int64(cur.Val))
  }
  return false
}
// @lc code=end

