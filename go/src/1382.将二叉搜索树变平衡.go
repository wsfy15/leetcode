/*
 * @lc app=leetcode.cn id=1382 lang=golang
 *
 * [1382] 将二叉搜索树变平衡
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
func balanceBST(root *TreeNode) *TreeNode {
	// 先获取升序序列，再构造二叉树
	var nums []int
	getSortNums(root, &nums)

	return generateBST(nums)
}

func generateBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  generateBST(nums[:mid]),
		Right: generateBST(nums[mid+1:]),
	}
}

func getSortNums(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	getSortNums(root.Left, nums)
	*nums = append(*nums, root.Val)
	getSortNums(root.Right, nums)
}

// @lc code=end

