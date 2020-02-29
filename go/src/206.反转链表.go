/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var dummy = new(ListNode)
	for head != nil {
		tmp := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = tmp
	}

	return dummy.Next
}

// @lc code=end
