/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = new(ListNode)
	dummy.Next = head

	var fast *ListNode = dummy
	var cur *ListNode = dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast != nil {
		for fast.Next != nil {
			fast = fast.Next
			cur = cur.Next
		}
	}

	cur.Next = cur.Next.Next
	return dummy.Next
}

// @lc code=end
