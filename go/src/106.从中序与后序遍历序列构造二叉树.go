/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
// 后序遍历最后一个元素是根节点
// 由于不会有重复元素，所以根据根节点的值，在中序遍历中可以得到左子树节点数量
// 然后分成左右两个子树 
func buildTree(inorder []int, postorder []int) *TreeNode {
  size := len(inorder)
  if size == 0 {
    return nil
  }
  root := &TreeNode{ // 后序遍历最后一个元素就是根节点
    Val: postorder[size-1],
  }
  
  i := 0
  for inorder[i] != root.Val {
    i++
  }
  // inorder[:i] 左子树
  // inorder[i+1:] 右子树
  root.Left = buildTree(inorder[:i], postorder[:i])
  root.Right = buildTree(inorder[i+1:], postorder[i:size-1])
  
  return root
}
// @lc code=end

