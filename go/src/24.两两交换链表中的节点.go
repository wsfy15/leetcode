/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var dummy ListNode 
	dummy.Next = head
	prev, cur := &dummy, dummy.Next

	for cur != nil && cur.Next != nil {
		prev.Next = cur.Next
		cur.Next= cur.Next.Next
		prev.Next.Next = cur
		prev = cur
		cur = cur.Next
	}

	return dummy.Next
}
// @lc code=end

