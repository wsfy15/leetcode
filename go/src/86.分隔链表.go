/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummyLeft, dummyRight := &ListNode{}, &ListNode{}
	left, right := dummyLeft, dummyRight

	for head != nil {
		if head.Val < x {
			left.Next = head
			head = head.Next
			left = left.Next
			left.Next = nil
		} else {
			right.Next = head
			head = head.Next
			right = right.Next
			right.Next = nil
		}
	}

	left.Next = dummyRight.Next
	return dummyLeft.Next
}
// @lc code=end

