/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
 type BSTIterator struct {
  inorder []int
  index int
}

func Constructor(root *TreeNode) BSTIterator {
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
  return BSTIterator{inorder: inorder, index: -1}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
  this.index++
  return this.inorder[this.index]
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
  return this.index + 1 < len(this.inorder)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

