/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	cur := root
	var parent *TreeNode
	for cur != nil && cur.Val != key {
		parent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		return root
	}

	// 有两个子节点，找到右子树中的最小值，替换当前节点
	if cur.Left != nil && cur.Right != nil {
		minNode := cur.Right
		minNodeParent := cur
		for minNode.Left != nil {
			minNodeParent = minNode
			minNode = minNode.Left
		}

		cur.Val = minNode.Val
		cur = minNode
		parent = minNodeParent
	}

	// cur只有一个子节点或没有子节点
	// 找到当前节点后继节点，作为parent节点的后继节点
	var child *TreeNode
	if cur.Left != nil {
		child = cur.Left
	} else if cur.Right != nil {
		child = cur.Right
	} else {
		child = nil
	}

	if parent == nil {
		root = child
	} else	if parent.Left == cur {
		parent.Left = child
	} else {
		parent.Right = child
	}

	return root
}
// @lc code=end

