/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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

 func sortedArrayToBST(nums []int) *TreeNode {
  if len(nums) == 0 {
    return nil
  } 
  mid := len(nums) >> 1
	root := &TreeNode{
    Val: nums[mid],
    Left: sortedArrayToBST(nums[:mid]),
    Right: sortedArrayToBST(nums[mid+1:]),
  }
  return root
}

 func sortedArrayToBST2(nums []int) *TreeNode {
  if len(nums) == 0 {
    return nil
  } 
  
  root := buildNode(nums, 0, len(nums) - 1)
  return root
}

func buildNode(nums []int, start, end int) *TreeNode {
  if start > end {
    return nil
  }
  
  mid := start + (end - start) >> 1
  root := &TreeNode{
    Val: nums[mid],
    Left: buildNode(nums, start, mid - 1),
    Right: buildNode(nums, mid + 1, end),
  }
  return root
}
// @lc code=end

