/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
 func isSymmetric(root *TreeNode) bool {
  if root == nil {
    return true
  }
  return dfs(root.Left, root.Right)
}

func dfs(t1, t2 *TreeNode) bool {
  if t1 == nil && t2 == nil {
    return true
  }
  if t1 == nil || t2 == nil {
    return false
  }
  
  return t1.Val == t2.Val && dfs(t1.Left, t2.Right) && dfs(t1.Right, t2.Left)
}

func isSymmetric2(root *TreeNode) bool {
  if root == nil {
    return true
  }
  
  queue := []*TreeNode{root}
  next := []*TreeNode{}
  nums := []int{}
  for len(queue) > 0 {
    cur := queue[0]
    queue = queue[1:]
    
    if cur == nil {
      nums = append(nums, 345345345)
      next = append(next, nil, nil)
    } else {
      nums = append(nums, cur.Val)
      next = append(next, cur.Left, cur.Right)
    }

    if len(queue) == 0 {
      queue, next = next, queue
      size := len(nums)
      for i, j := 0, size - 1; i < j; {
        if nums[i] != nums[j] {
          return false
        }
        i++
        j--
      }  
      nums = []int{}
      
      size = len(queue) - 1
      for size >= 0 {
        if queue[size] != nil {
          break
        }
        size--
      }
      if size == -1 {
        return true
      }
    }
    
  }
  
  return true
}
// @lc code=end

