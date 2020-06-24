/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
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
func sortedListToBST(head *ListNode) *TreeNode {
	// 快慢指针找中点
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{ Val: head.Val }
	}

	low, fast := head, head
	prev := &ListNode { Next: head }
	for fast != nil && fast.Next != nil {
		prev = prev.Next
		low = low.Next
		fast = fast.Next.Next
	}

	root := &TreeNode{ Val: low.Val }
	prev.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(low.Next)
	return root
}
// @lc code=end

