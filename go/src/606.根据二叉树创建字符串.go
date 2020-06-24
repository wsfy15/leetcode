/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
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
func tree2str(t *TreeNode) string {
	bytes := []byte{}
	preorder(t, &bytes)
	return string(bytes)
}

func preorder(t *TreeNode, bytes *[]byte) {
	if t == nil {
		return
	}

	*bytes = append(*bytes, []byte(strconv.Itoa(t.Val))...)
	if t.Left == nil && t.Right == nil {
		return
	}

	*bytes = append(*bytes, '(')
	preorder(t.Left, bytes)
	*bytes = append(*bytes, ')')

	if t.Right != nil {
		*bytes = append(*bytes, '(')
		preorder(t.Right, bytes)
		*bytes = append(*bytes, ')')
	}
}
// @lc code=end

