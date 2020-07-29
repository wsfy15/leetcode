/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	str1 := dfs(root1, "")
	str2 := dfs(root2, "")
	return str1 == str2
}

func dfs(root *TreeNode, cur string) string {
	if root == nil {
		return cur
	}
	if root.Left == nil && root.Right == nil {
		cur += strconv.Itoa(root.Val) + ","
		return cur
	}

	cur = dfs(root.Left, cur)
	cur = dfs(root.Right, cur)
	return cur
} 
// @lc code=end

