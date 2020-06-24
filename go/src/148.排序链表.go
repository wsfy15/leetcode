/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next:head}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		dummy = dummy.Next
	}

	dummy.Next = nil
	first, second := sortList(head), sortList(slow)
	return merge(first, second)
}

func merge(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			prev.Next = a
			a = a.Next
		} else {
			prev.Next = b
			b = b.Next
		}
		prev = prev.Next
		prev.Next = nil
	}
	if a != nil {
		prev.Next = a
	} else if b != nil {
		prev.Next = b
	}
	return dummy.Next
}
// @lc code=end

