/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	// reverse后半段链表,然后依次插入到前半段的每一个后面
	prev := &ListNode{ Next: head }
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		prev = prev.Next
	}

	if fast != nil {
		slow = slow.Next
		prev = prev.Next
	}

	prev.Next = nil
	first, second := head, reverse(slow)
	for second != nil {
		tmp := second.Next
		second.Next = first.Next
		first.Next = second
		second = tmp
		first = first.Next.Next
	}
}

func reverse(root *ListNode) *ListNode {
	dummy := &ListNode{}
	for root != nil {
		tmp := root.Next
		root.Next = dummy.Next
		dummy.Next = root
		root = tmp
	}

	return dummy.Next
}
// @lc code=end

