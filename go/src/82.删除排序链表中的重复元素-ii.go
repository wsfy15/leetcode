/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var dummy ListNode
	dummy.Next = head
	cur, prev := head, &dummy

	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			}
			cur = cur.Next
		} else {
			prev.Next = cur
			prev = cur
			cur = cur.Next
		}
	}

	prev.Next = nil
	return dummy.Next
}
// @lc code=end

