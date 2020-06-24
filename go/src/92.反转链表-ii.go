/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var dummy ListNode
	dummy.Next = head
	left, right := &dummy, &dummy
	for i := 0; i < n - m; i++ {
		right = right.Next
	}

	for i := 1; i < m; i++ {
		left, right = left.Next, right.Next
	}

	prev := left
	left, right = left.Next, right.Next

	for left != right {
		prev.Next = left.Next
		tmp := left.Next
		left.Next = right.Next
		right.Next = left
		left = tmp
	}

	return dummy.Next
}
// @lc code=end

