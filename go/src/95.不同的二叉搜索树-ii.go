/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func generateTrees(n int) []*TreeNode {
  if n == 0 {
    return []*TreeNode{}
  }
  return util(1, n)
}

func util(start, end int) []*TreeNode {
  if start > end {
    return []*TreeNode{nil,}
  }
  
  trees := make([]*TreeNode, 0)
  for i := start; i <= end; i++ {
    // 以i为根节点
    left := util(start, i - 1)
    right := util(i + 1, end)
    for ltree := range left {
      for rtree := range right {
        root := &TreeNode{
          Val: i,
          Left: left[ltree],
          Right: right[rtree],
        }
        trees = append(trees, root)
      }
    }
  }
  
  return trees
}
// @lc code=end

